package global_observer

import (
	"blogv2/observer/observer_article"
	"blogv2/observer/observer_comment"
)

var (
	ArticleNotifier *observer_article.ObserverArticleNotifier
	CommentNotifier *observer_comment.ObserverCommentNotifier
)
