package logic

import (
	"maxbell/dao/mysql"
	"maxbell/models"
	"maxbell/pkg/snowflake"
)

func SignUp(p *models.ParamSignup) (err error) {
	//判读用户是否存在
	if err = mysql.QueryUserbyusername(p.Username); err != nil {
		return err
	}
	//生成uid
	userId := snowflake.GenID()
	//构造user 实例
	user := models.User{
		Userid:   userId,
		Username: p.Username,
		Password: p.Password,
	}

	//存进数据库
	return mysql.InsertUser(&user)
}

func Login(p *models.ParamLogin) error {
	user := &models.User{

		Username: p.Username,
		Password: p.Password,
	}
	return mysql.Login(user)
}
