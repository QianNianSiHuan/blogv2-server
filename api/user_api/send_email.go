package user_api

import (
	"blogv2/common/res"
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	"blogv2/service/email_service"
	"blogv2/utils/email_store"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

type SendEmailRequest struct {
	Type  int8   `json:"type" binding:"oneof=1 2 3"` //1注册 2重置密码 3绑定邮箱
	Email string `json:"email" binding:"required"`
}
type SendEmailResponse struct {
	EmailID string `json:"emailID"`
}

func (UserApi) SendEmailView(c *gin.Context) {
	var cr SendEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	if !global.Config.Site.Login.EmailLogin {
		res.FailWithMsg(c, "站点未启用邮箱注册")
		return
	}
	code := base64Captcha.RandText(4, "0123456789")
	id := base64Captcha.RandomId()
	err = global.CaptchaStore.Set(id, code)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	switch cr.Type {
	case 1:
		var user models.UserModel
		err = global.DB.Take(&user, "email = ?", cr.Email).Error
		if err == nil {
			res.FailWithMsg(c, "该邮箱已存在")
			return
		}
		err = email_service.SendRegisterCode(cr.Email, code)
	case 2:
		var user models.UserModel
		err = global.DB.Take(&user, "email = ?", cr.Email).Error
		if err != nil {
			res.FailWithMsg(c, "该邮箱不存在")
			return
		}
		//还必须是邮箱注册
		//todo:该修改密码逻辑可后续优化
		if user.RegisterSource != enum.RegisterEmailSourceType {
			res.FailWithMsg(c, "非邮箱注册用户,不能更改密码")
			return
		}
		err = email_service.SendResetPwdCode(cr.Email, code)
	case 3:
		var user models.UserModel
		err = global.DB.Take(&user, "email = ?", cr.Email).Error
		if err == nil {
			res.FailWithMsg(c, "该邮箱已存在")
			return
		}
		err = email_service.SendBindEmailCode(cr.Email, code)
	}
	if err != nil {
		logrus.Errorf("邮件发送失败 %s", err)
		res.FailWithMsg(c, "邮件发送失败")
		return
	}
	global.EmailVerifyStore.Store(id, email_store.EmailStoreInfo{
		Email: cr.Email,
		Code:  code,
	})
	res.SuccessWithData(c, SendEmailResponse{
		EmailID: id,
	})
}
