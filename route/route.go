package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kashimashino/hackweek-group6/controller"
	"github.com/kashimashino/hackweek-group6/middlewares"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middlewares.CORS())  //加跨域中间件

	r.POST("user/registration", controller.UserRegistration) //注册接口
	r.POST("user/login", controller.UserLogin)               //登录接口

	userRouter := r.Group("/user")
	userRouter.Use(middlewares.JwtToken())
	{ //根据产品需求，暂定只有一个修改密码功能
		userRouter.PUT("/transformation", controller.ChangePw) //修改密码
	}

	tempRouter := r.Group("temp")
	tempRouter.Use(middlewares.JwtToken())
	{
		//获取图片
		tempRouter.GET("/:username", controller.Temperature)
		//获取sts token 供前端调用上传图片
		tempRouter.POST("/token", controller.StsToken)

	}

	if err := r.Run(":8080"); err != nil {
		panic("server启动失败")
	}
}
///
