package logic

import (
	"maxbell/dao/mysql"
	"maxbell/models"
	"maxbell/pkg/snowflake"
)

func SignUp(p *models.ParamSignup){
	//判读用户是否存在
	mysql.QueryUserbyusername()
	//生成uid
	snowflake.GenID()
	//密码加密

	//存进数据库
	mysql.InsertUser()
}