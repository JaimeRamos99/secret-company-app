package logic

import (
	json "encoding/json"
	fmt "fmt"

	database "github.com/JaimeRamos99/prueba-truora-2/database"
	structs "github.com/JaimeRamos99/prueba-truora-2/utils/structs"
	dgo "github.com/dgraph-io/dgo/v200"
)

func GetAllUsers(db *dgo.Dgraph) map[string]string {
	//call db package to query all the products
	res := database.GetAllUsers(db, false)
	//Parse the response to json
	res_json := fmt.Sprintf("%s\n", res.Json)

	//parse the json to struct
	var users *structs.Users
	json.Unmarshal([]byte(res_json), &users)

	//products map that contains every product and ensure uniqueness
	users_map := make(map[string]string)

	//store data into the map
	for _, usr := range users.Users {
		users_map[usr.UserId] = usr.Uid
	}
	return users_map

}

//function that found which of the users is a new one
func NewUsers(users_day structs.Users, users_db map[string]string) structs.Users {
	var new_users structs.Users
	map_users := make(map[string]int)

	for _, usr := range users_day.Users {
		//verify it's not on the db
		if _, ok := users_db[usr.UserId]; !ok {
			//only one user on the same day
			if _, ok2 := map_users[usr.UserId]; !ok2 {
				map_users[usr.UserId] = 1
				new_users.Users = append(new_users.Users, usr)
			}
		}
	}
	return new_users
}
