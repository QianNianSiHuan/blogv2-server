package user_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	jwts "blogv2/utils/jwt"
	"github.com/gin-gonic/gin"
	"time"
)

type UserDetailResponse struct {
	ID             uint                    `json:"id"`
	CreatedAt      time.Time               `json:"createdAt"`
	Username       string                  `json:"username"`
	Nickname       string                  `json:"nickname"`
	Avatar         string                  `json:"avatar"`
	Abstract       string                  `json:"abstract"`
	RegisterSource enum.RegisterSourceType `json:"registerSource"` // 注册来源
	CodeAge        int                     `json:"codeAge"`        // 码龄
	Role           enum.RoleType           `json:"role"`
	models.UserConfModel
	Email       string `json:"email"`
	UsePassword bool   `json:"usePassword"`
}

func (UserApi) UserDetailView(c *gin.Context) {
	claims := jwts.GetClaims(c)
	var user models.UserModel
	err := global.DB.Preload("UserConfModel").Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMsg(c, "用户不存在")
		return
	}

	var data = UserDetailResponse{
		ID:             user.ID,
		CreatedAt:      user.CreatedAt,
		Username:       user.Username,
		Nickname:       user.Nickname,
		Avatar:         user.Avatar,
		Abstract:       user.Abstract,
		RegisterSource: user.RegisterSource,
		Role:           user.Role,
		CodeAge:        user.CodeAge(),
		Email:          user.Email,
	}
	if user.Password != "" {
		data.UsePassword = true
	}

	if user.UserConfModel != nil {
		data.UserConfModel = *user.UserConfModel
	}
	res.SuccessWithData(c, data)
}
