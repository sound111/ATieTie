package logic

import (
	"TieTie/dao/myRedis"
	"TieTie/models"
	"strconv"
)

//direction =1
//直接点赞
//原本反对票，后来点赞

//该项目只允许 给一周内发布的帖子投票
//时间超过一周后，将redis中的数据存储到mysql中，删除redis中的相应存储

const perValue = 60 * 60 * 24 / 200

func PostVote(userId int64, p *models.ParamVoteData) (err error) {

	return myRedis.PostVote(strconv.Itoa(int(userId)), strconv.Itoa(int(p.PostId)), float64(p.Direction*perValue))
}
