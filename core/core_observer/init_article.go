package core_observer

import (
	"blogv2/global/global_observer"
	"blogv2/observer/observer_article"
	"github.com/sirupsen/logrus"
)

func initArticleObserver() {
	logrus.Info("文章观察者加载中...")
	articleLookHistory := observer_article.NewArticleLookHistory()
	articleDigg := observer_article.NewArticleDigg()
	articleCollect := observer_article.NewArticleCollect()
	articleComment := observer_article.NewArticleComment()
	articleExamine := observer_article.NewArticleExamine()

	global_observer.ArticleNotifier = observer_article.NewArticleNotifier()

	global_observer.ArticleNotifier.AddArticleLookListener(articleLookHistory)
	global_observer.ArticleNotifier.AddArticleDiggIncrListener(articleDigg)
	global_observer.ArticleNotifier.AddArticleDiggDecListener(articleDigg)
	global_observer.ArticleNotifier.AddArticleCollectDecListener(articleCollect)
	global_observer.ArticleNotifier.AddArticleCollectIncrListener(articleCollect)
	global_observer.ArticleNotifier.AddArticleCommentIncrListener(articleComment)
	global_observer.ArticleNotifier.AddArticleCommentDecListener(articleComment)
	global_observer.ArticleNotifier.AddArticleExamineListener(articleExamine)
	logrus.Info("文章观察者加载完成")
	return
}
