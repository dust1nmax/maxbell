package routes

import (
	"maxbell/controllers"
	"maxbell/logger"
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

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	return r
}
