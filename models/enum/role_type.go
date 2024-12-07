package enum

type RoleType int8

const (
	AdminRole   RoleType = 1 //管理员
	UserRole    RoleType = 2 //用户
	VisitorRole RoleType = 3 //访客
)
