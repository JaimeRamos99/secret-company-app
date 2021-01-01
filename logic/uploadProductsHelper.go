package logic

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
	strings "strings"

	database "github.com/JaimeRamos99/prueba-truora-2/database"
	structs "github.com/JaimeRamos99/prueba-truora-2/utils/structs"
	dgo "github.com/dgraph-io/dgo/v200"
)

//parse array of bytes that contains the products of the given day to an array os structs
func ProductsDecoder(bytes []byte) []structs.Product {

	//parse de data in bytes format to string
	//and split each transaction info by \n
	input_str := string(bytes)
	products_array_string := strings.Split(input_str, "\n")

	//for each product string, the attributes are applited by '
	//then a struct is created with those attrs
	var prods_array []structs.Product
	for _, pr := range products_array_string {
		//The last line of the response is empty
		if len(pr) > 0 {
			attr_prods := strings.Split(pr, "'")
			product_id := attr_prods[0]
			product_name := attr_prods[1]
			price, errorr := strconv.Atoi(attr_prods[2])
			if errorr != nil {
				price = 0
			}
			//create a product struct with the attrs
			product := *structs.NewProduct(product_id, product_name, price)
			prods_array = append(prods_array, product)
		}
	}
	return prods_array
}

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
