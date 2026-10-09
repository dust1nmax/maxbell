package routes

import (
	"maxbell/controllers"
	"maxbell/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.Ginlogger(), logger.GinRecovery(true))

	//注册业务路由
	r.POST("/signup", controllers.SignUpHandler)

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	return r
}
