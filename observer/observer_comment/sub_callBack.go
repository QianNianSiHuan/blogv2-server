package observer_comment

import "blogv2/service/redis_service/redis_comment"

type CommentSub struct {
}

func NewCommentSub() *CommentSub {
	return &CommentSub{}
}

func (c CommentSub) AfterCommentSubDec(commentID uint, n int) {
	redis_comment.SetCacheApply(commentID, n)
}

func (c CommentSub) AfterCommentSubIncr(commentID uint) {
	redis_comment.SetCacheApply(commentID, 1)
}
