package logic

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"

	cache "github.com/JaimeRamos99/prueba-truora-2/cache"
	database "github.com/JaimeRamos99/prueba-truora-2/database"
	structs "github.com/JaimeRamos99/prueba-truora-2/utils/structs"
	dgo "github.com/dgraph-io/dgo/v200"
	redis "github.com/go-redis/redis/v8"
)

//Main function that handles the whole process of uploading the data of a given day
func UploadData(db *dgo.Dgraph, rdb *redis.Client, date string, unix_time int64) bool {
	//check the cache to find if the date has been uploaded before
	data_in_cache := cache.CheckDate(rdb, date)

	//the date is not in the cache
	if !data_in_cache {
		//check the database to know if the data of that day has been uploaded
		//and converts the db response into a manipulable format
		res := database.CheckDate(db, date)
		res_string := fmt.Sprintf("%s\n", res.Json)
		var resp *structs.RespArray
		json.Unmarshal([]byte(res_string), &resp)

		//The data has to be uploaded, because it's not in the db
		if len(resp.Query) == 0 {
			unix_time_string := strconv.FormatInt(unix_time, 10)
			usr_map := UploadUsers(db, unix_time_string)
			prods_map := UploadProducts(db, unix_time_string)
			upload_status := HandleTransactions(db, unix_time_string, usr_map, prods_map)
			//check if all the steps in the upload process were succesfull
			if upload_status {
				cache.SetDate(rdb, date)
				database.AddUploadDate(db, date)
				//data was uploaded succesfully
				return true
			}
		}
		//data was in the db
		return false
	}
	//the date is in the cache, so had already been uploaded
	return false
}
