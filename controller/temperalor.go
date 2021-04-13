package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kashimashino/hackweek-group6/model"
	"github.com/kashimashino/hackweek-group6/service"
	"github.com/kashimashino/hackweek-group6/utils"
	"net/http"
	"strconv"
)

func Temperature(c *gin.Context) {
	var tpa service.TempCalcService
	var err error
	if tpa.Temperature, err = strconv.Atoi(c.Param("temperature")); err == nil {
		info := tpa.TempCalc()

		if err := info.Msg; err != utils.GetErrMsg(200) {
			c.JSON(http.StatusBadRequest, err)
		} else {
			fmt.Println(info.Data)
			c.JSON(http.StatusOK, &model.Response{
				Status: 200,
				Msg:    utils.GetErrMsg(200),
				Data:   info.Data,
			})
		}
	}
}
