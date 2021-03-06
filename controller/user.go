package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kashimashino/hackweek-group6/middlewares"
	"github.com/kashimashino/hackweek-group6/model"
	"github.com/kashimashino/hackweek-group6/service"
	"github.com/kashimashino/hackweek-group6/utils"
	"net/http"
)

// UserRegistration 管理用户注册
func UserRegistration(c *gin.Context) {
	var srv service.UserRegisterService
	if err := c.ShouldBind(&srv); err == nil { //绑定JSON
		if user, err := srv.Register(); err != nil {
			c.JSON(http.StatusBadRequest, err) //输出error
		} else {
			token := middlewares.SetToken(user.Username) //生成token
			c.JSON(http.StatusOK, &model.Response{
				Data:   user,
				Status: 200,
				Msg:    utils.GetErrMsg(200),
				Token:  token,
			})
		}
	}
}

//UserLogin 管理用户登录
func UserLogin(c *gin.Context) {
	var srv service.UserLoginService
	if err := c.ShouldBind(&srv); err == nil {
		if user, err := srv.Login(); err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			token := middlewares.SetToken(user.Username) //生成token
			c.JSON(http.StatusOK, &model.Response{
				Data:   user,
				Status: 200,
				Msg:    utils.GetErrMsg(200),
				Token:  token,
			})
		}
	}
}

//ChangePw 用户修改密码
func ChangePw(c *gin.Context) {
	var srv service.ChangePwService
	if err := c.ShouldBind(&srv); err == nil {
		fmt.Println("------------------------")
		if err := srv.ChangePW(); err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			c.JSON(http.StatusOK, &model.Response{
				Status: 200,
				Msg:    utils.GetErrMsg(200),
			})
		}
	} else {
		fmt.Println("-------------------------")
	}
}

//StsToken 拿sts token
func StsToken(c *gin.Context) {
	var srv service.OssService
	if data := srv.UploadToken(); data != nil {
		c.JSON(http.StatusOK, data)
	} else {
		c.JSON(http.StatusBadRequest, data)
	}
}
