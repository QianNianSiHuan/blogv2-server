package models

import (
	"blogv2/global"
	"blogv2/models/ctype"
	"blogv2/models/enum"
	"blogv2/service/redis_service/redis_article"
	"blogv2/service/text_service"
	_ "embed"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sync"
)

type ArticleModel struct {
	Model
	Title         string             `gorm:"size:32" json:"title"`
	Abstract      string             `gorm:"size:256" json:"abstract"`
	Content       string             `json:"content"`
	CategoryID    *uint              `json:"categoryID"` //分类
	CategoryModel *CategoryModel     `gorm:"foreignKey:CategoryID" json:"-"`
	TagList       ctype.List         `gorm:"type:longtext" json:"tagList" `
	Cover         string             `gorm:"size:256" json:"cover"`
	UserID        uint               `json:"userID"`
	UserModel     UserModel          `gorm:"foreignKey:UserID" json:"-"`
	LookCount     int                `json:"lookCount"`
	DiggCount     int                `json:"diggCount"`
	CommentCount  int                `json:"commentCount"`
	CollectCount  int                `json:"collectCount"`
	OpenComment   bool               `json:"openComment"` //评论开关
	Status        enum.ArticleStatus `json:"status"`      //状态
}

//go:embed mappings/article_mapping.json
var articleMappings string

func (a ArticleModel) Mapping() string {
	return articleMappings
}
func (a ArticleModel) Index() string {
	return "article_index"
}

func (a *ArticleModel) BeforeDelete(tx *gorm.DB) (err error) {
	// 评论
	var commentList []CommentModel
	global.DB.Find(&commentList, "article_id = ?", a.ID).Delete(&commentList)
	// 点赞
	var diggList []ArticleDiggModel
	global.DB.Find(&diggList, "article_id = ?", a.ID).Delete(&diggList)
	// 收藏
	var collectList []UserArticleCollectModel
	global.DB.Find(&collectList, "article_id = ?", a.ID).Delete(&collectList)
	// 置顶
	var topList []UserTopArticleModel
	global.DB.Find(&topList, "article_id = ?", a.ID).Delete(&topList)
	// 浏览
	var lookList []UserArticleLookHistoryModel
	global.DB.Find(&lookList, "article_id = ?", a.ID).Delete(&lookList)

	logrus.Infof("删除关联评论 %d 条", len(commentList))
	logrus.Infof("删除关联点赞 %d 条", len(diggList))
	logrus.Infof("删除关联收藏 %d 条", len(collectList))
	logrus.Infof("删除关联置顶 %d 条", len(topList))
	logrus.Infof("删除关联浏览 %d 条", len(lookList))
	return
}
func (a *ArticleModel) AfterCreate(tx *gorm.DB) (err error) {
	// 创建文章之后的钩子函数
	// 只有发布中的文章会放到全文搜索里面去
	var wg sync.WaitGroup
	if a.Status != enum.ArticleStatusPublished {
		return nil
	}
	go func() {
		wg.Add(1)
		err = a.setTextSearchIndex()
		if err != nil {
			logrus.Error(err.Error())
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		a.setArticleSearchIndex()
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		a.setTagSearchIndex()
		wg.Done()
	}()
	wg.Wait()
	return nil
}

func (a *ArticleModel) AfterDelete(tx *gorm.DB) (err error) {
	// 删除之后
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		var textList []TextModel
		tx.Find(&textList, "article_id = ?", a.ID)
		if len(textList) > 0 {
			logrus.Infof("删除全文记录 %d", len(textList))
			tx.Delete(&textList)
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		redis_article.RemoveTagAgg(a.ID, a.TagList...)
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		text_service.DeleteArticleParticiple(a.ID)
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		text_service.DeleteTextParticiple(a.ID)
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		redis_article.ClearArticleSortByID(a.ID)
		wg.Done()
	}()
	wg.Wait()
	return nil
}

func (a *ArticleModel) AfterUpdate(tx *gorm.DB) (err error) {
	// 正文发生了变化，才去做转换
	a.AfterDelete(tx)
	a.AfterCreate(tx)
	return nil
}

func (a *ArticleModel) setArticleSearchIndex() {
	cacheCollect := redis_article.GetAllCacheCollect(1)
	cacheComment := redis_article.GetAllCacheComment(1)
	cacheDigg := redis_article.GetAllCacheDigg(1)
	cacheLook := redis_article.GetAllCacheLook(1)
	collectCount := cacheCollect[a.ID] + a.CollectCount
	commentCount := cacheComment[a.ID] + a.CommentCount
	diggCount := cacheDigg[a.ID] + a.DiggCount
	lookCount := cacheLook[a.ID] + a.LookCount
	redis_article.SetCacheCollectSortByCount(a.ID, collectCount)
	redis_article.SetCacheCommentSortByCount(a.ID, commentCount)
	redis_article.SetCacheDiggSortByCount(a.ID, diggCount)
	redis_article.SetCacheLookSortByCount(a.ID, lookCount)
	redis_article.SetCacheAllSort(a.ID)
	words := text_service.TextParticiple(a.Title, a.Abstract)
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		redis_article.SetArticleSearchIndex(a.ID, words)
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		redis_article.SetArticleSearchWords(a.ID, words)
		wg.Done()
	}()
	wg.Wait()
}

func (a *ArticleModel) setTextSearchIndex() error {
	var list []TextModel
	var textParticipleList []text_service.ParticipleTextModel

	textList := text_service.MdContentTransformation(a.ID, a.Title, a.Content)
	if len(textList) == 0 {
		return nil
	}
	for _, model := range textList {
		list = append(list, TextModel{
			ArticleID: model.ArticleID,
			Head:      model.Head,
			Body:      model.Body,
		})
	}
	err := global.DB.Create(&list).Error
	if err != nil {
		logrus.Error(err)
		return err
	}
	for _, _list := range list {
		textParticiple := text_service.ParticipleTextModel{
			ID: _list.ID,
			TextModel: text_service.TextModel{
				ArticleID: _list.ArticleID,
				Head:      _list.Head,
				Body:      _list.Body,
			},
		}
		textParticipleList = append(textParticipleList, textParticiple)
	}

	//redis分词索引
	text_service.TextSearchParticiple(textParticipleList...)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return nil
}
func (a *ArticleModel) setTagSearchIndex() {
	for _, tag := range a.TagList {
		if tag == "" {
			continue
		}
		redis_article.SetTagAgg(tag, a.ID)
		redis_article.SetTagAggAdd(tag)
	}
}
