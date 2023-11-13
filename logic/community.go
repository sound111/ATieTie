package logic

import (
	"TieTie/dao/mysql"
	"TieTie/models"
)

func GetCommunityList() ([]*models.Community, error) {
	//查询数据库
	return mysql.GetCommunityList()
}

func GetCommunityInfo(id int64) (*models.Community, error) {
	return mysql.GetCommunityInfo(id)
}
