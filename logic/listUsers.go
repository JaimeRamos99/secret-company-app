package logic

import (
	json "encoding/json"
	fmt "fmt"

	database "github.com/JaimeRamos99/prueba-truora-2/database"
	structs "github.com/JaimeRamos99/prueba-truora-2/utils/structs"
	dgo "github.com/dgraph-io/dgo/v200"
)

func ListAllUsers(db *dgo.Dgraph) *structs.Users {
	resp := database.GetAllUsers(db, true)
	//Parse the response to json
	res_json := fmt.Sprintf("%s\n", resp.Json)

	//parse the json to struct
	var users *structs.Users
	json.Unmarshal([]byte(res_json), &users)
	return users
}
