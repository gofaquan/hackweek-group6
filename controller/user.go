package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kashimashino/hackweek-group6/model"
	"github.com/kashimashino/hackweek-group6/service"
	"github.com/kashimashino/hackweek-group6/utils"
	"net/http"
)

//管理用户注册
func UserRegistration(c *gin.Context) {
	var srv service.UserRegisterService
	if err := c.ShouldBind(&srv); err == nil { //绑定JSON
		if user, err := srv.Register(); err != nil {
			c.JSON(http.StatusBadRequest, err) //输出error
		} else {
			c.JSON(http.StatusOK, &model.Response{
				Data:   user,
				Status: 200,
				Msg:    utils.GetErrMsg(200),
			})
		}
	}
}

func UserLogin(c *gin.Context) {
	var srv service.UserLoginService
	if err := c.ShouldBind(&srv); err == nil {
		if user, err := srv.Login(); err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			c.JSON(http.StatusOK, &model.Response{
				Data:   user,
				Status: 200,
				Msg:    utils.GetErrMsg(200),
			})
		}
	}
}

func ChangePw(c *gin.Context) {
	var srv service.ChangePwService
	if err := c.ShouldBind(&srv); err == nil {
		if err := srv.ChangePW(); err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			c.JSON(http.StatusOK, &model.Response{
				Status: 200,
				Msg:    utils.GetErrMsg(200),
			})
		}
	}
}
