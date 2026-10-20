package controllers

import (
	"maxbell/logic"
	"maxbell/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreatePostHandler
func CreatePostHandler(c *gin.Context) {
	// 参数校验与处理
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("c.ShouldBindJSON(p) failed", zap.Error(err))
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

// GetPostHandler
func GetPostHandler(c *gin.Context) {
	// 获取参数 （url里的帖子id）
	idstr := c.Param("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		zap.L().Error("get post detial failed with invalid param ", zap.Error(err))
		ReponseError(c, CodeInvalidParams)
		return
	}
	// 通过id 取出帖子数据 （查数据库）
	data, err := logic.GetPostDetial(id)
	if err != nil {
		zap.L().Error("logic.GetPostDetial failed", zap.Error(err))
		ReponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ReponseSuccess(c, data)

}

// GetPostListHandler 获取帖子列表的接口
func GetPostListHandler(c *gin.Context) {

	page, size := GetPageInfo(c)

	// 获取数据
	datas, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ReponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ReponseSuccess(c, datas)
}


