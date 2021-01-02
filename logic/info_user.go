package logic

import (
	json "encoding/json"
	fmt "fmt"

	database "github.com/JaimeRamos99/prueba-truora-2/database"
	structs "github.com/JaimeRamos99/prueba-truora-2/utils/structs"
	dgo "github.com/dgraph-io/dgo/v200"
)

func UserInfo(db *dgo.Dgraph, userId string) {
	//parse the bd response to struct
	resp_general_info := database.User_general_info(db, userId)
	general_info_json := fmt.Sprintf("%s\n", resp_general_info.Json)

	//create a struct for the user general info
	var general_info *structs.InfoArray
	json.Unmarshal([]byte(general_info_json), &general_info)
	fmt.Println(general_info_json)
	fmt.Println(general_info)
	//resp_same_ips := database.Same_ips(db, userId)
	//resp_recommendation := database.Recommendation(db)
}
