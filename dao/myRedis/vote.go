package myRedis

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
)

const oneWeekSeconds = 7 * 24 * 60 * 60

var ErrPostTimeExpire = errors.New("投票时间已过")

func PostVote(userId, postId string, value float64) (err error) {
	postTime := db.ZScore(getRedisKey(KeyPostTimeZSet), postId).Val()

	if float64(time.Now().Unix())-postTime > oneWeekSeconds {
		err = ErrPostTimeExpire
		return
	}

	//当前用户给当前帖子的投票记录
	ov := db.ZScore(getRedisKey(KeyVoteZSetPrefix+postId), userId).Val()

	pipeline := db.TxPipeline()

	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), value-ov, postId)

	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyVoteZSetPrefix+postId), userId)
	} else {
		pipeline.ZAdd(getRedisKey(KeyVoteZSetPrefix+postId), redis.Z{
			Score:  value,
			Member: userId,
		})
	}

	_, err = pipeline.Exec()
	return
}
