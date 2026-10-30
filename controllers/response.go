package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*

{
	"code"
	"msg"
	"data"
}

*/

type ResponseData struct {
	Code Codemsg `json:"code"`
	Msg  any     `json:"msg"`
	Data any     `json:"data,omitempty"`
}

func ReponseError(c *gin.Context, code Codemsg) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func ReponseSuccess(c *gin.Context, data any) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}

func ReponseErrorwithMsg(c *gin.Context, code Codemsg, msg any) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}