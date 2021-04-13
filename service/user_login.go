package service

import (
	"github.com/kashimashino/hackweek-group6/model"
	"github.com/kashimashino/hackweek-group6/utils"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	Username string `form:"username" json:"username" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// Login 用户登录函数
func (service *UserLoginService) Login() (*model.User, *model.Response) {
	user := model.User{
		Username: service.Username,
		Password: service.Password,
	}
	if err := CRUD.UserLogin(&user); err != nil {
		return nil, &model.Response{
			Status: 1003,
			Msg:    utils.GetErrMsg(1003), //用户不存在
		}
	}

	if user.CheckPassword(service.Password) != nil {
		return nil, &model.Response{
			Status: 1013,
			Msg:    utils.GetErrMsg(1013), //帐号或密码错误
		}
	}
	return &user, nil
}
