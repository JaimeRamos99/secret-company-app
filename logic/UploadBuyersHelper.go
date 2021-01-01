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
	res := database.GetAllUsers(db)
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

func NewUsers(prods_day []structs.Product, prods_db map[string]string) []structs.Product {
	var new_prods []structs.Product
	for _, prod := range prods_day {
		if _, ok := prods_db[prod.ProductID]; !ok {
			new_prods = append(new_prods, prod)
		}
	}
	return new_prods
}
