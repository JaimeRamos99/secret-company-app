package logic

import (
	fmt "fmt"
	ioutil "io/ioutil"
	log "log"
	http "net/http"
	strconv "strconv"
	strings "strings"

	utils "github.com/JaimeRamos99/prueba-truora-2/utils"
	structs "github.com/JaimeRamos99/prueba-truora-2/utils/structs"
	dgo "github.com/dgraph-io/dgo/v200"
)

func UploadProducts(db *dgo.Dgraph, date string) bool {

	//ctx := context.Background()

	//creating the request to fetch the resource
	url_string := fmt.Sprintf(utils.Base_url, "products", date)
	req, er := http.NewRequest("GET", url_string, nil)
	if er != nil {
		log.Fatal(er)
		return false
	}

	//Commit the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer resp.Body.Close()

	//read the body response (our transactions data) and store as bytes
	bytes, erro := ioutil.ReadAll(resp.Body)
	if erro != nil {
		log.Fatalln(erro)
		return false
	}

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
	GetAllProducts(db)
	/*mu := &api.Mutation{
		CommitNow: true,
	}

	products_json, error := json.Marshal(prods_array)
	if error != nil {
		log.Fatal(err)
		return false
	}
	mu.SetJson = products_json
	assigned, err := db.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}*/
	return true
}
