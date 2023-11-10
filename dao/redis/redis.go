package redis

import (
	"Web/settings"
	"strconv"

	"github.com/go-redis/redis"
)

var db *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {

	db = redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + strconv.Itoa(cfg.Port),
		Password: cfg.Pwd, // no password set
		DB:       cfg.DB,  // use default db
	})

	//ctx := context.Background()
	_, err = db.Ping().Result()
	if err != nil {
		return
	}

	return
}

func Close() (err error) {
	err = db.Close()
	return
}
