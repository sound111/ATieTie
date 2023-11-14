package myRedis

// myRedis key 使用冒号分隔，形成命名空间
const (
	KeyPrefix         = "bluebell:" //公共前缀
	KeyPostTimeZSet   = "post:time"
	KeyPostScoreZSet  = "post:score"
	KeyVoteZSetPrefix = "post:vote:" //参数是post_id
)

// 加上前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}
