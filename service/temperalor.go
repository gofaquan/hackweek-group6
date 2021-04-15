package service

import (
	"github.com/kashimashino/hackweek-group6/model"
)

type TempCalcService struct {
	Username string `json:"username"`
}

func (tpa *TempCalcService) TempCalc() ([]model.Temperalor, *model.Response) {
	tempColor := &model.Temperalor{
		Username: tpa.Username,
	}
	info := CRUD.CalcTemp(tempColor)
	if info == nil {
		return nil, &model.Response{
			Status: 500,
			Msg:    "你还没有发布过图片呢",
		}
	}
	return *info, nil
}
