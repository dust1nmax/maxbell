package routes

import (
	"maxbell/controllers"
	"maxbell/logger"
	"maxbell/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	//gin.SetMode(gin.ReleaseMode) //设置为发布模式
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.Ginlogger(), logger.GinRecovery(true))

	//注册业务路由
	r.POST("/signup", controllers.SignUpHandler)
	//登陆业务
	r.POST("/login", controllers.LoginHandler)

	r.GET("/ping", middlewares.JWTAuthMiddleware(), func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ping")
	})

	return r
}

