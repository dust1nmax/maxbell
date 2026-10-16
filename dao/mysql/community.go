package mysql

import (
	"database/sql"
	"maxbell/models"

	"go.uber.org/zap"
)

func GetCommunityList() (data []*models.Community, err error) {
	sqlstr := "select community_id, community_name from community"
	if err := db.Select(&data, sqlstr); err != nil {
		if err == sql.ErrNoRows{
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}
