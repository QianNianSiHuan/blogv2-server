package observer_article

import (
	"github.com/gin-gonic/gin"
)

type AfterArticleLookListener interface {
	AfterArticleLook(c *gin.Context, articleID uint) //观察者得到通知后要触发的动作
}
type AfterArticleDiggIncrListener interface {
	AfterArticleDiggIncr(articleID uint) //观察者得到通知后要触发的动作
}
type AfterArticleDiggDecListener interface {
	AfterArticleDiggDec(articleID uint) //观察者得到通知后要触发的动作
}

type AfterArticleExamineListener interface {
	AfterArticleExamine(articleID uint, status int8) //观察者得到通知后要触发的动作
}
type AfterArticleCollectIncrListener interface {
	AfterArticleCollectIncr(articleID uint, status int8) //观察者得到通知后要触发的动作
}
type AfterArticleCollectDecListener interface {
	AfterArticleCollectDec(articleID uint, status int8) //观察者得到通知后要触发的动作
}
type AfterArticleCommentIncrListener interface {
	AfterArticleCommentIncr(articleID uint) //观察者得到通知后要触发的动作
}
type AfterArticleCommentDecListener interface {
	AfterArticleCommentDec(articleID uint, status int) //观察者得到通知后要触发的动作
}
type ArticleNotifier interface {
	AddArticleLookListener(...AfterArticleLookListener)
	AddArticleDiggIncrListener(...AfterArticleDiggIncrListener)
	AddArticleDiggDecListener(...AfterArticleDiggDecListener)
	AddArticleExamineListener(...AfterArticleExamineListener)
	AddArticleCollectIncrListener(...AfterArticleCollectIncrListener)
	AddArticleCollectDecListener(...AfterArticleCollectDecListener)
	AddArticleCommentIncrListener(...AfterArticleCommentIncrListener)
	AddArticleCommentDecListener(...AfterArticleCommentDecListener)

	RemoveArticleLookListener(...AfterArticleLookListener)
	RemoveArticleDiggIncrListener(...AfterArticleDiggIncrListener)
	RemoveArticleDiggDecListener(...AfterArticleDiggDecListener)
	RemoveArticleExamineListener(...AfterArticleExamineListener)
	RemoveArticleCollectIncrListener(...AfterArticleCollectIncrListener)
	RemoveArticleCollectDecListener(...AfterArticleCollectDecListener)
	RemoveArticleCommentIncrListener(...AfterArticleCommentIncrListener)
	RemoveArticleCommentDecListener(...AfterArticleCommentDecListener)

	AfterArticleLookNotify(*gin.Context, uint)
	AfterArticleDiggIncrNotify(uint)
	AfterArticleDiggDecNotify(uint)
	AfterArticleExamineNotify(uint, int8)
	AfterArticleCollectIncrNotify(uint)
	AfterArticleCollectDecNotify(uint)
	AfterArticleCommentIncrNotify(uint)
	AfterArticleCommentDecNotify(uint, int)
}
