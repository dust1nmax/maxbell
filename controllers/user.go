package controllers

import (
	"fmt"
	"maxbell/logic"
	"maxbell/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	//1. 参数处理
	var p models.ParamSignup
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误
		zap.L().Error("sign up with invalid params", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}

	//手动对参数进行详细的校验
	if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
		zap.L().Error("sign up with invalid params")
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}

	fmt.Println(p)
	//2. 业务处理
	logic.SignUp(&p)
	//3. 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
