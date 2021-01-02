package logic

import (
	json "encoding/json"
	fmt "fmt"

	database "github.com/JaimeRamos99/prueba-truora-2/database"
	structs "github.com/JaimeRamos99/prueba-truora-2/utils/structs"
	dgo "github.com/dgraph-io/dgo/v200"
)

func UserInfo(db *dgo.Dgraph, userId string) {
	//parse the bd response for general info to struct
	resp_general_info := database.User_general_info(db, userId)
	general_info_json := fmt.Sprintf("%s\n", resp_general_info.Json)

	//create a struct for the user general info
	var general_info *structs.InfoArray
	json.Unmarshal([]byte(general_info_json), &general_info)

	//
	resp_same_ips := database.Same_ips(db, userId)
	same_ips_json := fmt.Sprintf("%s\n", resp_same_ips.Json)

	//
	var same_ips *structs.DataSameIps
	json.Unmarshal([]byte(same_ips_json), &same_ips)
	fmt.Println(same_ips)
	//resp_recommendation := database.Recommendation(db)
}
