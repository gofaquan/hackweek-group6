package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kashimashino/hackweek-group6/controller"
)

func InitRouter() {
	r := gin.Default()

	userRouter := r.Group("/user")
	{
		userRouter.POST("/registration", controller.UserRegistration) //注册接口
		userRouter.POST("/login", controller.UserLogin)               //登录接口
		userRouter.PUT("/transformation", controller.ChangePw)        //修改密码
	}

	r.Run(":8080")
}
