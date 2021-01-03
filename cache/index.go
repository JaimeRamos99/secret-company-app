package cache

import (
	context "context"
	log "log"

	redis "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func NewClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	log.Print("Established connection with redis")
	return rdb
}

func CheckDate(rdb *redis.Client, key string) bool {
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
	err := rdb.Set(ctx, date, true, 0).Err()
	if err != nil {
		panic(err)
	}
}
