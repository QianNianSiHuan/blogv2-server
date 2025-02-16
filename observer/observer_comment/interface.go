package observer_comment

type AfterCommentDiggIncrListener interface {
	AfterCommentDiggIncr(articleID uint) //观察者得到通知后要触发的动作
}
type AfterCommentDiggDecListener interface {
	AfterCommentDiggDec(articleID uint) //观察者得到通知后要触发的动作
}

type AfterCommentSubIncrListener interface {
	AfterCommentSubIncr(articleID uint) //观察者得到通知后要触发的动作
}
type AfterCommentSubDecListener interface {
	AfterCommentSubDec(commentID uint, n int) //观察者得到通知后要触发的动作
}
type CommentNotifier interface {
	AddCommentDiggIncrListener(...AfterCommentDiggIncrListener)
	AddCommentDiggDecListener(...AfterCommentDiggDecListener)
	AddCommentSubIncrListener(...AfterCommentSubIncrListener)
	AddCommentSubDecListener(...AfterCommentSubDecListener)

	RemoveCommentDiggIncrListener(...AfterCommentDiggIncrListener)
	RemoveCommentDiggDecListener(...AfterCommentDiggDecListener)
	RemoveCommentSubIncrListener(...AfterCommentSubIncrListener)
	RemoveCommentSubDecListener(...AfterCommentSubDecListener)

	AfterCommentDiggIncrNotify(uint)
	AfterCommentDiggDecNotify(uint)
	AfterCommentSubIncrNotify(uint)
	AfterCommentSubDecNotify(uint, int)
}
