package service

import (
	"github.com/kashimashino/hackweek-group6/model"
	"github.com/kashimashino/hackweek-group6/utils"
)

type TempCalcService struct {
	Temperature int `json:"temperature"`
}

func (temp *TempCalcService) TempCalc() (string, *model.Response) {
	tempColor := &model.Temperalor{
		Temperature: temp.Temperature,
	}
	info := CRUD.CalcTemp(tempColor)
	if info == nil {
		return "", &model.Response{
			Status: 500,
			Msg:    "获取图片失败",
		}
	}
	return "", &model.Response{
		Status: 200,
		Msg:    utils.GetErrMsg(200),
		Data:   info.Url,
	}
}
