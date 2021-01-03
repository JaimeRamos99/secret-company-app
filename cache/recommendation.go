package cache

import (
	context "context"
	json "encoding/json"
	fmt "fmt"
	time "time"

	structs "github.com/JaimeRamos99/prueba-truora-2/utils/structs"
	redis "github.com/go-redis/redis/v8"
)

func GetRecommendation(rdb *redis.Client) {
	ctx := context.Background()
	recommendation_json, err := rdb.Get(ctx, "recommendation").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("no está")
		}
		panic(err)
	}
	fmt.Println(recommendation_json)
}

func SetRecommendation(rdb *redis.Client, recommendation []structs.TopProduct) {
	ctx := context.Background()
	//save the recommendation struct as an []byte
	recommendation_json, err := json.Marshal(recommendation)
	if err != nil {
		panic(err)
	}

	er := rdb.Set(ctx, "recommendation", recommendation_json, 24*time.Hour).Err()
	if er != nil {
		panic(er)
	}

}
