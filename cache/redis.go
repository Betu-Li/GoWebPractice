package cache

import (
	"GoWebPractice/setting"
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var (
	Rdb  *redis.Client
	Rctx context.Context
)

func Init() {
	if setting.Conf.RedisConfig == nil {
		log.Fatal("Redis configuration is not initialized")
	}

	Rctx = context.Background()

	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: setting.Conf.RedisConfig.RedisPassword,
		DB:       setting.Conf.RedisConfig.RedisDB,
	})

	_, err := Rdb.Ping(Rctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}

func Zscore(id, score int) redis.Z {
	return redis.Z{Score: float64(score), Member: id}
}
