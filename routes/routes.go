package routes

import (
	"maxbell/controllers"
	"maxbell/logger"
	"maxbell/middlewares"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	//gin.SetMode(gin.ReleaseMode) //设置为发布模式
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.Ginlogger(), logger.GinRecovery(true))

	v1:= r.Group("/api/v1")

	//注册业务路由
	v1.POST("/signup", controllers.SignUpHandler)
	//登陆业务
	v1.POST("/login", controllers.LoginHandler)

	v1.Use(middlewares.JWTAuthMiddleware()) //应用JWT认证中间件

	{
		v1.GET("/community", controllers.CommunityHandler)
		v1.GET("/community/:id", controllers.CommunityDetailHandler)

		v1.POST("/post", controllers.CreatePostHandler)
	}

	return r
}

