package controllers

import (
	"maxbell/logic"

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


