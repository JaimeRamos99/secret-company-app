package logic

import (
	json "encoding/json"
	fmt "fmt"

	cache "github.com/JaimeRamos99/prueba-truora-2/cache"
	database "github.com/JaimeRamos99/prueba-truora-2/database"
	structs "github.com/JaimeRamos99/prueba-truora-2/utils/structs"
	dgo "github.com/dgraph-io/dgo/v200"
	redis "github.com/go-redis/redis/v8"
)

func UserInfo(db *dgo.Dgraph, rdb *redis.Client, userId string) *structs.AllUserInfo {

	//parse the bd response for other users using the same ip of a given user
	resp_general_info := database.User_general_info(db, userId)
	general_info_json := fmt.Sprintf("%s\n", resp_general_info.Json)

	//create a struct for the user general info
	var general_info *structs.InfoArray
	json.Unmarshal([]byte(general_info_json), &general_info)

	//parse the bd response for other users using the same ip of a given user
	resp_same_ips := database.Same_ips(db, userId)
	same_ips_json := fmt.Sprintf("%s\n", resp_same_ips.Json)

	//create a struct for the users that share the same ips of a given user
	var same_ips *structs.DataSameIps
	json.Unmarshal([]byte(same_ips_json), &same_ips)

	//parse the bd response of the recommendation to struct
	resp_recommendation := database.Recommendation(db)
	recommendation_json := fmt.Sprintf("%s\n", resp_recommendation.Json)

	//create a struct for the recommedation response
	var recommendations *structs.Result
	json.Unmarshal([]byte(recommendation_json), &recommendations)

	//top three products
	recommended_products := ThreeBestSellers(recommendations)
	cache.SetRecommendation(rdb, recommended_products)
	cache.GetRecommendation(rdb)

	//all the info asked in the test together
	all_info := structs.NewAllUserInfo(general_info, same_ips, recommended_products)
	return all_info
}
