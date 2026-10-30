package controllers

import (
	"maxbell/logic"
	"maxbell/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func PostVoteHandler(c *gin.Context) {
	// 参数校验
	p := new(models.ParamPostVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) //类型断言
		if !ok {
			ReponseError(c, CodeInvalidParams)
			return
		}
		errdata := removeTopStruct(errs.Translate(trans)) //翻译并去掉错误信息中最外层的结构体名称
		ReponseErrorwithMsg(c, CodeInvalidParams, errdata)
		return
	}
	// 业务逻辑

	userid, err := GetCurrentUser(c)
	if err != nil{
		ReponseError(c, CodeInvalidParams)
		return
	}
	if err := logic.PostVote(userid, p); err!=nil{
		zap.L().Error("logic.PostVote(userid, p) failed", zap.Error(err))
		ReponseError(c, CodeServerBusy)
		return
	}
	ReponseSuccess(c, nil)
}
