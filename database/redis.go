package database

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func InitRedis(ctx context.Context, address string, password string, db string) *redis.Client {
	var i, err = strconv.Atoi(db)
	if err != nil {
		panic(err)
	}

	var redisdb = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       i,
	})

	err = redisdb.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}

	return redisdb
}
