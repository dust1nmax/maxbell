package controllers

type Codemsg int64

const (
	CodeSuccess Codemsg = 1000 + iota
	CodeInvalidParams
	CodeuserExist
	CodeuserNotExist
	CodeInvalidPassword
	CodeServerBusy
)

var codeMsgMap = map[Codemsg]string{
	CodeSuccess:         "success",
	CodeInvalidParams:   "请求参数错误",
	CodeuserExist:       "用户存在",
	CodeuserNotExist:    "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",
}

func (c Codemsg) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}

	return msg
}
