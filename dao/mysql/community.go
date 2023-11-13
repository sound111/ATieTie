package mysql

import (
	"TieTie/models"
	"TieTie/myError"
	"database/sql"
	"errors"

	"go.uber.org/zap"
)

func GetCommunityList() (communities []*models.Community, err error) {
	sqlStr := "select community_id,community_name from community"

	err = db.Select(&communities, sqlStr)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			zap.L().Warn("no community data in db")
			err = nil
		} else {
			return
		}
	}

	return
}

func GetCommunityInfo(id int64) (m *models.Community, err error) {
	m = new(models.Community)
	sqlStr := "select community_id,community_name,introduction,create_time from community where community_id=?"
	err = db.Get(m, sqlStr, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = myError.ErrorInvalidID
		} else {
			return
		}
	}

	return
}
