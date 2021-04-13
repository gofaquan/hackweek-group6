package service

import (
	"fmt"
	"github.com/kashimashino/hackweek-group6/model"
	"github.com/kashimashino/hackweek-group6/utils"
)

type TempCalcService struct {
	Url         string `json:"url"`
	Temperature int    `json:"temperature"`
}

func (temp *TempCalcService) TempCalc() *model.Response {
	tempColor := &model.Temperalor{
		Temperature: temp.Temperature,
	}
	info := CRUD.CalcTemp(tempColor)
	if info == nil {
		return &model.Response{
			Status: 500,
			Msg:    "获取图片失败",
		}
	}
	fmt.Println(info.Url)
	fmt.Println("----------------------")
	return &model.Response{
		Status: 200,
		Msg:    utils.GetErrMsg(200),
		Data:   info.Url,
	}
}
