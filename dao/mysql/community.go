package mysql

import (
	"database/sql"
	"maxbell/models"

	"go.uber.org/zap"
)

func GetCommunityList() (data []*models.Community, err error) {
	sqlstr := "select community_id, community_name from community"
	if err := db.Select(&data, sqlstr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

func GetCommunityDetialbyid(id int64) (community *models.CommunityDetail,err error) {
	community = new(models.CommunityDetail)
	sqlstr := "select community_id, community_name, introduction, create_time from community where community_id = ?"
	 
	if err := db.Get(community, sqlstr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidId
		} 
	}
	return community, err
}
