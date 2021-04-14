package controller

import (
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

		if url, err := tpa.TempCalc(); err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			c.JSON(http.StatusOK, &model.Response{
				Status: 200,
				Msg:    utils.GetErrMsg(200),
				Data:   url,
			})
		}
	}
}

func PostTemp(c *gin.Context) {
	service := service.UploadTokenService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Post()
		c.JSON(200, res)
	} else {
		c.JSON(200, err)
	}

}
