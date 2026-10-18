package controllers

import (
	"maxbell/logic"
	"maxbell/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context) {
	// 参数校验与处理
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		ReponseError(c, CodeInvalidParams)
		return
	}
	// 从c 获取当前用户id
	userID, err := GetCurrentUser(c)
	if err != nil {
		ReponseError(c, CodeNeedAuth)
		return
	}
	p.AuthorID = userID
	//业务实现
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ReponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ReponseSuccess(c, CodeSuccess)
}
