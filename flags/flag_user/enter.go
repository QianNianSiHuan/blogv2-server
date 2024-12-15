package flag_user

import (
	"blogv2/global"
	"blogv2/models"
	"blogv2/models/enum"
	"blogv2/unitls/pwd"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

type FlagUser struct {
}

func (FlagUser) Creat() {
	var role enum.RoleType
	var err error
	for !(role == 1 || role == 2 || role == 3) {
		fmt.Println("选择角色 1超级管理员 2普通用户 3访客")
		_, err = fmt.Scan(&role)
		if err != nil {
			logrus.Errorf("输入错误 %s", err)
			return
		}
	}
	var username string
	for {
		fmt.Println("请输入用户名:")
		_, err = fmt.Scan(&username)
		if err != nil || username == "" {
			logrus.Errorf("输入错误 %s", err)
			continue
		}
		var model models.UserModel
		err = global.DB.Take(&model, "username = ?", username).Error
		if err == nil {
			logrus.Errorf("用户名已存在")
			continue
		}
		break
	}
	var password, rePassword []byte
	for {
		fmt.Println("请输入密码:")
		password, err = terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			logrus.Errorf("输入错误 %s", err)
			continue
		}

		fmt.Println("请再次输入密码:")
		rePassword, err = terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			logrus.Errorf("输入错误 %s", err)
			continue
		}
		if string(password) != string(rePassword) {
			fmt.Println("两次密码不一致")
			continue
		}
		break
	}
	hashPwd, _ := pwd.GenerateFromPassword(string(password))
	err = global.DB.Create(&models.UserModel{
		Username:       username,
		Nickname:       "",
		RegisterSource: enum.RegisterTerminalSourceType,
		Password:       hashPwd,
		Role:           role,
	}).Error
	if err != nil {
		logrus.Errorf(err.Error())
		return
	}
	logrus.Info("创建成功")
}
