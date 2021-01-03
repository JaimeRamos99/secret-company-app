package cache

import (
	log "log"
	"os/exec"

	utils "github.com/JaimeRamos99/prueba-truora-2/utils"
	redis "github.com/go-redis/redis/v8"
)

func NewClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	log.Print("Established connection with redis")
	return rdb
}

func DeleteAllPreviousKeys() {
	cmd := exec.Command("/bin/sh", "-c", utils.CommandDeleteRedisData)
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
	log.Print("Command Successfully Executed to delete data in redis: ", cmd)
}
