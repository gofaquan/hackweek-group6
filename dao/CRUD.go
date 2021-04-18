package dao

import (
	"fmt"
	"github.com/kashimashino/hackweek-group6/model"
)

type CRUD struct {
}

func (CRUD *CRUD) UserRegisterService(count int64, username string) int64 {
	DB.Model(&model.User{}).Where("username = ?", username).Count(&count)
	return count
}

func (CRUD *CRUD) CreateUser(user *model.User) error {
	err := DB.Create(&user).Error
	return err
}

func (CRUD *CRUD) UserLogin(user *model.User) error {
	err := DB.Where("username = ?", user.Username).First(&user).Error
	return err
}

func (CRUD *CRUD) ChangePw(user *model.User, newPw string) error {
	err := DB.Model(&model.User{}).Where("username = ?", user.Username).Update("password", newPw).Error
	return err
}

func (CRUD *CRUD) CalcTemp(tempColor *model.Temperalor) (*[]model.Temperalor, int64) {
	var infos []model.Temperalor
	var count int64
	DB.Model(&model.Temperalor{}).Where("username = ?", tempColor.Username).Find(&infos).Count(&count)
	fmt.Println(infos)
	return &infos, count
}
