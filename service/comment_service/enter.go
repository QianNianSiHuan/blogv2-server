package comment_service

import (
	"blogv2/global"
	"blogv2/models"
	"blogv2/service/redis_service/redis_comment"
	"blogv2/service/text_service"
	"time"
)

// GetRootComment 获取一个评论的根评论
func GetRootComment(commentID uint) (model *models.CommentModel) {
	var comment models.CommentModel
	err := global.DB.Take(&comment, commentID).Error
	if err != nil {
		return nil
	}
	if comment.ParentID == nil {
		// 没有父评论了，那么他就是根评论
		return &comment
	}
	return GetRootComment(*comment.ParentID)
}

// GetParents 获取一个评论的所有父评论
func GetParents(commentID uint) (list []models.CommentModel) {
	var comment models.CommentModel
	err := global.DB.Take(&comment, commentID).Error
	if err != nil {
		return
	}
	list = append(list, comment)
	if comment.ParentID != nil {
		// 没有父评论了，那么他就是根评论
		list = append(list, GetParents(*comment.ParentID)...)
	}
	return
}

// GetCommentTree 获取评论树
func GetCommentTree(model *models.CommentModel) {
	global.DB.Preload("SubCommentList").Take(model)
	for _, commentModel := range model.SubCommentList {
		GetCommentTree(commentModel)
	}
}

// GetCommentTreeV2 获取评论树
func GetCommentTreeV2(id uint) (model *models.CommentModel) {
	model = &models.CommentModel{
		Model: models.Model{ID: id},
	}

	global.DB.Preload("SubCommentList").Take(model)
	for i := 0; i < len(model.SubCommentList); i++ {
		commentModel := model.SubCommentList[i]
		item := GetCommentTreeV2(commentModel.ID)
		model.SubCommentList[i] = item
	}
	return
}

type CommentResponse struct {
	ID           uint               `json:"id"`
	CreatedAt    time.Time          `json:"createdAt"`
	Content      string             `json:"content"`
	UserID       uint               `json:"userID"`
	UserNickname string             `json:"userNickname"`
	UserAvatar   string             `json:"userAvatar"`
	ArticleID    uint               `json:"articleID"`
	ParentID     *uint              `json:"parentID"`
	DiggCount    int                `json:"diggCount"`
	ApplyCount   int                `json:"applyCount"`
	SubComments  []*CommentResponse `json:"subComments"`
	IsDigg       bool               `json:"isDigg"`
}

func GetCommentTreeV4(id uint, userDiggMap map[uint]bool) (res *CommentResponse) {
	return getCommentTreeV4(id, 1, userDiggMap)
}

func getCommentTreeV4(id uint, line int, userDiggMap map[uint]bool) (res *CommentResponse) {
	model := &models.CommentModel{
		Model: models.Model{ID: id},
	}

	global.DB.Preload("UserModel").Preload("SubCommentList").Take(model)
	model.Content = text_service.ReplaceSensitiveWords(model.Content, "*")
	res = &CommentResponse{
		ID:           model.ID,
		CreatedAt:    model.CreatedAt,
		Content:      model.Content,
		UserID:       model.UserID,
		UserNickname: model.UserModel.Nickname,
		UserAvatar:   model.UserModel.Avatar,
		ArticleID:    model.ArticleID,
		ParentID:     model.ParentID,
		DiggCount:    model.DiggCount + redis_comment.GetCacheDigg(model.ID),
		ApplyCount:   redis_comment.GetCacheApply(model.ID),
		SubComments:  make([]*CommentResponse, 0),
		IsDigg:       userDiggMap[model.ID],
	}
	if line >= global.Config.Site.Article.CommentLine {
		return
	}
	for _, commentModel := range model.SubCommentList {
		res.SubComments = append(res.SubComments, getCommentTreeV4(commentModel.ID, line+1, userDiggMap))
	}
	return
}

// GetCommentOneDimensional 评论一维化
func GetCommentOneDimensional(id uint) (list []models.CommentModel) {
	model := models.CommentModel{
		Model: models.Model{ID: id},
	}

	global.DB.Preload("SubCommentList").Take(&model)
	list = append(list, model)
	for _, commentModel := range model.SubCommentList {
		subList := GetCommentOneDimensional(commentModel.ID)
		list = append(list, subList...)
	}
	return
}
