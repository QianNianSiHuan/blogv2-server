package core_observer

import (
	"blogv2/global/global_observer"
	"blogv2/observer/observer_comment"
	"github.com/sirupsen/logrus"
)

func initCommentObserver() {
	logrus.Info("文章评论观察者加载中...")
	commentDigg := observer_comment.NewCommentDigg()
	commentSub := observer_comment.NewCommentSub()

	global_observer.CommentNotifier = observer_comment.NewCommentNotifier()

	global_observer.CommentNotifier.AddCommentDiggIncrListener(commentDigg)
	global_observer.CommentNotifier.AddCommentDiggDecListener(commentDigg)
	global_observer.CommentNotifier.AddCommentSubIncrListener(commentSub)
	global_observer.CommentNotifier.AddCommentSubDecListener(commentSub)
	logrus.Info("文章评论观察者加载完成")
	return
}
