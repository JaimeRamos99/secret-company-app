package cache

import (
	context "context"
	json "encoding/json"
	time "time"

	structs "github.com/JaimeRamos99/prueba-truora-2/utils/structs"
	redis "github.com/go-redis/redis/v8"
)

func GetRecommendation(rdb *redis.Client) *[]structs.TopProduct {
	ctx := context.Background()
	recommendation_json, err := rdb.Get(ctx, "recommendation").Result()
	if err != nil {
		if err == redis.Nil {
			return nil
		}
		panic(err)
	}
	var recommendation_struct *[]structs.TopProduct
	json.Unmarshal([]byte(recommendation_json), &recommendation_struct)
	return recommendation_struct
}

func SetRecommendation(rdb *redis.Client, recommendation []structs.TopProduct) {
	ctx := context.Background()
	//save the recommendation struct as an []byte
	recommendation_json, err := json.Marshal(recommendation)
	if err != nil {
		panic(err)
	}
	//cache the recommendation of the day for 24 hours
	er := rdb.Set(ctx, "recommendation", recommendation_json, 24*time.Hour).Err()
	if er != nil {
		panic(er)
	}

}
