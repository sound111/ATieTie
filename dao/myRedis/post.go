package myRedis

import (
	"TieTie/models"
	"time"

	"github.com/go-redis/redis"
)

func CreatePost(p *models.Post) (err error) {
	pipeline := db.TxPipeline() //事务，要么同时成功，要么同时失败

	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: p.PostId,
	})

	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: p.PostId,
	})

	_, err = pipeline.Exec()
	return
}
