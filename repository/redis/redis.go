package redis

import (
	"apihut-server/config"
	"errors"
	"fmt"

	"github.com/go-redis/redis"
)

var (
	ErrValueNotExit = errors.New("值不存在")
)

var rdb *redis.Client

func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Conf.RedisConfig.Host, config.Conf.RedisConfig.Port),
		Password: "",
		DB:       0,
		PoolSize: 100,
	})

	if _, err = rdb.Ping().Result(); err != nil {
		return err
	}

	fmt.Println("Redis ready!")

	return nil
}

func Close() {
	_ = rdb.Close()
}
