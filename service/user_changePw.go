package service

import (
	"github.com/kashimashino/hackweek-group6/model"
	"github.com/kashimashino/hackweek-group6/utils"
	"golang.org/x/crypto/bcrypt"
)

type ChangePwService struct {
	User         model.User `json:"user"`
	Password     string
	NewPw        string `json:"new_pw" form:"new_pw" binding:"required,min=8,max=40"`
	ConfirmNewPw string `json:"confirm_new_pw" form:"confirm_new_pw" binding:"required,min=8,max=40"`
}

func (service *ChangePwService) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 1)
	if err != nil {
		return err
	}
	service.NewPw = string(bytes)
	return nil
}

func (service *ChangePwService) ChangePW() *model.Response {
	if err := service.User.CheckPassword(service.Password); err != nil {
		return &model.Response{
			Status: 1002,
			Msg:    utils.GetErrMsg(1002), //密码错误
		}
	}

	if service.NewPw != service.ConfirmNewPw {
		return &model.Response{
			Status: 1010,
			Msg:    utils.GetErrMsg(1010), //两次输入的密码不一致
		}
	}

	if err := service.SetPassword(service.NewPw); err != nil {
		return &model.Response{
			Status: 1011,
			Msg:    utils.GetErrMsg(1011), //密码加密失败
		}
	}

	if err := CRUD.ChangePw(&service.User, service.NewPw); err != nil {
		return &model.Response{
			Status: 1014,
			Msg:    utils.GetErrMsg(1014),
		}
	}

	return nil
}
