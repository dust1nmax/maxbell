package controllers

import (
	"maxbell/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ---- 社区 ----

func CommunityHandler(c *gin.Context) {
	//查询所有社区（community_id, name）以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList failed", zap.Error(err))
		ReponseError(c, CodeServerBusy)
		return
	}

	ReponseSuccess(c, data)
}

func CommunityDetailHandler(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		ReponseError(c, CodeInvalidParams)
		return
	}	
	//查询所有社区（community_id, name）以列表的形式返回
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityList failed", zap.Error(err))
		ReponseError(c, CodeServerBusy)
		return
	}

	ReponseSuccess(c, data)
}
