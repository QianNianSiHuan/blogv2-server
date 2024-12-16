package user_server

import "blogv2/models"

type UserService struct {
	userModel models.UserModel
}

func NewUserServiceApp(user models.UserModel) *UserService {
	return &UserService{
		userModel: user,
	}
}
