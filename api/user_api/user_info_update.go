package user_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	jwts "blogv2/unitls/jwt"
	"blogv2/unitls/maps"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

type UserInfoUpdateRequest struct {
	Username    *string   `json:"username" s-u:"username"`
	Nickname    *string   `json:"nickname" s-u:"nickname"`
	Avatar      *string   `json:"avatar" s-u:"avatar"`
	Abstract    *string   `json:"abstract" s-u:"abstract"`
	LikeTags    *[]string `json:"likeTags" s-u-c:"like_tags"`
	OpenCollect *bool     `json:"openCollect" s-u-c:"open_collect"`
	OpenFollow  *bool     `json:"openFollow" s-u-c:"open_follow"`
	OpenFans    *bool     `json:"openFans" s-u-c:"open_fans"`
	HomeStyleID *uint     `json:"homeStyleID" s-u-c:"home_style_id"`
}

func (UserApi) UserInfoUpdateView(c *gin.Context) {
	var cr UserInfoUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	userMap, err := maps.StructToMap(cr, "s-u")
	if err != nil {
		res.FailWithMsg(c, "用户map转化失败")
		return
	}
	logrus.Info(userMap)
	userConfMap, err := maps.StructToMap(cr, "s-u-c")
	if err != nil {
		res.SuccessWithMsg(c, "用户配置表转map失败")
		return
	}

	claims := jwts.GetClaims(c)

	if len(userMap) > 0 {
		var userModel models.UserModel
		err = global.DB.Preload("UserConfModel").
			Take(&userModel, claims.UserID).Error
		if err != nil {
			logrus.Info(err)
			res.FailWithMsg(c, "用户不存在")
			return
		}
		logrus.Info(userModel.UserConfModel)
		//判断
		if cr.Username != nil {
			var count int64
			global.DB.Model(models.UserModel{}).
				Where("username = ? and id <> ?", *cr.Username, claims.UserID).
				Count(&count)
			if count > 1 {
				res.FailWithMsg(c, "用户名已经被使用")
				return
			}
			if *cr.Username != userModel.Username {
				if userModel.UserConfModel.UpdateUsernameTime != nil {
					if time.Now().Sub(*userModel.UserConfModel.UpdateUsernameTime).Hours() < 720 {
						res.FailWithMsg(c, "30天内只能修改一次")
						return
					}
					userConfMap["update_username_time"] = time.Now()
				}
			}
		}
		if cr.Nickname != nil || cr.Avatar != nil {
			if userModel.RegisterSource == enum.RegisterQQSourceType {
				res.FailWithMsg(c, "QQ注册的账号不能修改昵称和头像")
				return
			}
		}
		err = global.DB.Model(&userModel).Updates(userMap).Error
		if err != nil {
			res.FailWithMsg(c, "用户信息修改失败")
			return
		}
	}
	//用户配置的修改
	if len(userConfMap) > 0 {
		var userConfModel models.UserConfModel
		err = global.DB.Take(&userConfModel, "user_id", claims.UserID).Error
		if err != nil {
			res.FailWithMsg(c, "用户不存在")
			return
		}
		err = global.DB.Model(&userConfModel).Updates(userConfModel).Error
		if err != nil {
			res.FailWithMsg(c, "用户信息修改失败")
			return
		}
	}

	res.SuccessWithMsg(c, "用户信息更新成功")
}
