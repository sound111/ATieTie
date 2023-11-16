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

func GetPostsInOrder(p *models.ParamPostList) ([]string, error) {
	var key string
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	} else if p.Order == models.OrderTime {
		key = getRedisKey(KeyPostTimeZSet)
	}

	start := p.Size * (p.Page - 1)
	end := start + p.Size - 1

	return db.ZRevRange(key, start, end).Result()
}

func GetPostVoteData(postIds []string) (data []int64, err error) {
	data = make([]int64, 0, len(postIds))

	//使用pipeline一次发送多条命令，减少向redis发送命令的次数，减少RTT
	pipeline := db.Pipeline()

	for _, id := range postIds {
		key := getRedisKey(KeyVoteZSetPrefix + id)
		pipeline.ZCard(key)
	}

	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}

	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}

	return
}
