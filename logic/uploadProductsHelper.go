package logic

import (
	json "encoding/json"
	fmt "fmt"

	database "github.com/JaimeRamos99/prueba-truora-2/database"
	structs "github.com/JaimeRamos99/prueba-truora-2/utils/structs"
	dgo "github.com/dgraph-io/dgo/v200"
)

//get all products from the db
func GetAllProducts(db *dgo.Dgraph) map[string]string {
	//call db package to query all the products
	res := database.GetAllProducts(db)
	//Parse the response to json
	res_json := fmt.Sprintf("%s\n", res.Json)

	//parse the json to struct
	var products *structs.Products
	json.Unmarshal([]byte(res_json), &products)

	//products map that contains every product and ensure uniqueness
	products_map := make(map[string]string)

	//store data into the map
	for _, prod := range products.Products {
		products_map[prod.ProductId] = prod.Uid
	}
	return products_map
}

//determine which products are new to upload
func NewProducts(prods_day []structs.Product, prods_db map[string]string) []structs.Product {
	var new_prods []structs.Product
	//loops through products of the day and store the ones that are not on the db... new ones
	for _, prod := range prods_day {
		if _, ok := prods_db[prod.ProductID]; !ok {
			new_prods = append(new_prods, prod)
		}
	}
	return new_prods
}
