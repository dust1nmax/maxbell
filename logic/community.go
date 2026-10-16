package logic

import (
	"maxbell/dao/mysql"
	"maxbell/models"
)

func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()

}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetialbyid(id)
}
