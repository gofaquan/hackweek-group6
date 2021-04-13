package service

import (
	"github.com/kashimashino/hackweek-group6/dao"
	"github.com/kashimashino/hackweek-group6/model"
	"github.com/kashimashino/hackweek-group6/utils"
)

var count int64
var CRUD dao.CRUD

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Username        string `form:"username" json:"username" binding:"required,min=5,max=18"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=16"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" binding:"required,min=8,max=16"`
}

// Valid 验证表单
func (service *UserRegisterService) Valid() *model.Response {
	//验证密码是否输入一致
	if service.ConfirmPassword != service.Password {
		return &model.Response{
			Status: 1010,
			Msg:    utils.GetErrMsg(1010),
		}
	}

	//若count>0，则说明用户已存在
	count = CRUD.UserRegisterService(0, service.Username)
	if count > 0 {
		return &model.Response{
			Status: 1001,
			Msg:    utils.GetErrMsg(1001),
		}
	}

	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() (*model.User, *model.Response) {
	user := model.User{
		Username: service.Username,
		Password: service.Password,
	}

	// 表单验证
	if err := service.Valid(); err != nil {
		return nil, err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return nil, &model.Response{
			Status: 1011,
			Msg:    utils.GetErrMsg(1011),
		}
	}

	// 创建用户
	if err := CRUD.CreateUser(&user); err != nil {
		return nil, &model.Response{
			Status: 1012,
			Msg:    utils.GetErrMsg(1012),
		}
	}

	return &user, nil
}
