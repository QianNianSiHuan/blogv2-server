package enum

type CommentStatus int8

const (
	CommentStatusExamine   CommentStatus = 1 //审核中
	CommentStatusPublished CommentStatus = 2 //已发布
	CommentStatusFail      CommentStatus = 3 //已拒绝
)
