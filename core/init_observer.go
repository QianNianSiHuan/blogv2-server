package core

import (
	"blogv2/global/global_observer"
	"blogv2/observer/observer_article"
	"github.com/sirupsen/logrus"
)

func InitObserver() {
	articleLookHistory := observer_article.NewArticleLookHistory()
	global_observer.AfterDetailNotifier = observer_article.NewAfterArticleDetailNotifier()
	global_observer.AfterDetailNotifier.AddListener(articleLookHistory)
	logrus.Info("观察者加载完成")
	return
}
