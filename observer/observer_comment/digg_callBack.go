package observer_comment

import "blogv2/service/redis_service/redis_comment"

type CommentDigg struct{}

func NewCommentDigg() *CommentDigg {
	return &CommentDigg{}
}

func (c CommentDigg) AfterCommentDiggIncr(articleID uint) {
	redis_comment.SetCacheDigg(articleID, 1)
}

func (c CommentDigg) AfterCommentDiggDec(articleID uint) {
	redis_comment.SetCacheDigg(articleID, -1)
}
