package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"maxbell/models"
)

var secret = "max"

// 把每一步数据库操作封装成函数等到logic 层调用
func QueryUserbyusername(username string) (err error) {
	sqlstr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlstr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

func InsertUser(user *models.User) (err error) {
	user.Password = encryptPassword(user.Password)
	sqlstr := `insert into user(user_id, username, password) values(?,?,?)`
	_, err = db.Exec(sqlstr, user.Userid, user.Username, user.Password)
	return

}

// 密码加密
func encryptPassword(opassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(opassword)))
}

func Login(user *models.User) (err error) {
	opassword := user.Password
	sqlstr := `select user_id, username, password from user where username = ?`
	err = db.Get(user, sqlstr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return nil
	}
	password := encryptPassword(opassword)
	if password != user.Password {
		return ErrorInvalidParams
	}
	return
}
