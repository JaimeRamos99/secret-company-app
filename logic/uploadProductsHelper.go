package logic

import (
	json "encoding/json"
	fmt "fmt"

	database "github.com/JaimeRamos99/prueba-truora-2/database"
	structs "github.com/JaimeRamos99/prueba-truora-2/utils/structs"
	dgo "github.com/dgraph-io/dgo/v200"
)

func GetAllProducts(db *dgo.Dgraph) {
	//call db package to query all the products
	res := database.GetAllProducts(db)
	//Parse the response to json
	res_json := fmt.Sprintf("%s\n", res.Json)

	//parse the json to struct
	var resp *structs.Products
	json.Unmarshal([]byte(res_json), &resp)
	fmt.Println(resp.Products[0].Uid)
}
