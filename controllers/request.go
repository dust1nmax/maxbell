package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var ContextUserIdKey = "userId"

var ErrorUserNotLogin = errors.New("用户登陆失败")

func GetCurrentUser(c *gin.Context) (userID int64,err error) {
	uid, ok := c.Get(ContextUserIdKey)
	if !ok {
		err = ErrorUserNotLogin
		return 
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return 
	}
	return
}
