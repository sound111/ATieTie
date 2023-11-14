package logic

import (
	"TieTie/dao/myRedis"
	"TieTie/dao/mysql"
	"TieTie/models"
)

func CreatePost(p *models.Post) (err error) {
	err = mysql.CreatePost(p)
	if err != nil {
		return
	}

	err = myRedis.CreatePost(p)
	return
}

func GetPostInfo(id int64) (p *models.ParamPostInfo, err error) {
	return mysql.GetPostInfo(id)
}

func GetPostList() (p []*models.Post, err error) {
	return mysql.GetPostList()
}
