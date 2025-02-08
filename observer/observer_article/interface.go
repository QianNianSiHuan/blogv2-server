package observer_article

import (
	"github.com/gin-gonic/gin"
)

type AfterArticleDetailSuccessListener interface {
	AfterArticleDetailSuccess(c *gin.Context, articleID uint) //观察者得到通知后要触发的动作
}

type Notifier interface {
	AddListener(...AfterArticleDetailSuccessListener)
	RemoveListener(...AfterArticleDetailSuccessListener)
	Notify(*gin.Context, uint)
}
