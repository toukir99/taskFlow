package db

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "taskFlow-redis",
		DB:       0,
	})

	_, err := Rdb.Ping(context.Background()).Result()
	if err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}
}

func GetRedisClient() *redis.Client {
	return Rdb
}
