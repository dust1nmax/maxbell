package logic

import (
	"maxbell/dao/mysql"
	"maxbell/models"
)

func GetCommunityList() ([]*models.Community,	error) {
	return  mysql.GetCommunityList()

}
