package controllers

import (
	"errors"
	"fmt"
	"maxbell/dao/mysql"
	"maxbell/logic"
	"maxbell/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	//1. 参数处理
	var p models.ParamSignup
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误
		zap.L().Error("sign up with invalid params", zap.Error(err))

		//判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ReponseError(c, CodeInvalidParams)
			return
		}
		ReponseErrorwithMsg(c, CodeInvalidParams, removeTopStruct(errs.Translate(trans)))
		return
	}

	//手动对参数进行详细的校验
	// if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
	// 	zap.L().Error("sign up with invalid params")
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "请求参数有误",
	// 	})
	// 	return
	// }

	fmt.Println(p)
	//2. 业务处理
	if err := logic.SignUp(&p); err != nil {
		if errors.Is(err, mysql.ErrorUserExist){
			ReponseError(c, CodeuserExist)
			return
		}
		ReponseError(c, CodeServerBusy)
		return
	}
	//3. 返回响应
	ReponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context){
	//获取校验参数
	p := new(models.ParamLogin)
	if err:= c.ShouldBindJSON(p); err!= nil{
		//请求参数有误
		zap.L().Error("login with invalid params", zap.Error(err))

		//判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ReponseError(c, CodeInvalidParams)
			return
		}
		ReponseErrorwithMsg(c, CodeInvalidParams, removeTopStruct(errs.Translate(trans)))
		return
	}
	//业务处理
	if err := logic.Login(p); err!= nil{
		zap.L().Error("login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist){
			ReponseError(c, CodeuserNotExist)
			return
		}
		ReponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ReponseSuccess(c, nil)
}