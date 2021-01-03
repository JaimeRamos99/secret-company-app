package cache

import (
	context "context"

	redis "github.com/go-redis/redis/v8"
)

func CheckDate(rdb *redis.Client, key string) bool {
	ctx := context.Background()
	_, err := rdb.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return false
		}
		panic(err)

	}
	return true
}

func SetDate(rdb *redis.Client, date string) {
	ctx := context.Background()
	err := rdb.Set(ctx, date, true, 0).Err()
	if err != nil {
		panic(err)
	}
}
