package service

import (
	"github.com/kashimashino/hackweek-group6/model"
	"github.com/kashimashino/hackweek-group6/utils"
)

type PostTemp struct {
	Username string `json:"username"`
	Time     string `json:"time"`
	Avatar   string `json:"avatar"`
}

// PostTemp 上传图片
func (service *PostTemp) PostTemp() *model.Response {
	pp := model.PostPicture{
		Username: service.Username,
	}

	if err := CRUD.PostTemp(&pp); err != nil {
		return &model.Response{
			Status: 1015,
			Msg:    utils.GetErrMsg(1015), //上传失败
		}
	}

	return &model.Response{}
}
