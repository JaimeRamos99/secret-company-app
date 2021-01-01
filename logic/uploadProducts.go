package logic

import (
	context "context"
	json "encoding/json"
	fmt "fmt"
	ioutil "io/ioutil"
	log "log"
	http "net/http"

	utils "github.com/JaimeRamos99/prueba-truora-2/utils"
	dgo "github.com/dgraph-io/dgo/v200"
	api "github.com/dgraph-io/dgo/v200/protos/api"
)

func UploadProducts(db *dgo.Dgraph, date string) map[string]string {

	ctx := context.Background()

	//creating the request to fetch the resource
	url_string := fmt.Sprintf(utils.Base_url, "products", date)
	req, er := http.NewRequest("GET", url_string, nil)
	if er != nil {
		log.Fatal(er)
	}

	//Commit the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	//read the body response (our transactions data) and store as bytes
	bytes, erro := ioutil.ReadAll(resp.Body)
	if erro != nil {
		log.Fatalln(erro)
	}

	//receivs an stuc array
	prods_array := ProductsDecoder(bytes)

	//get all the products in the db and determine
	//which of the products of the given day
	all_products_db := GetAllProducts(db)
	new_prods := NewProducts(prods_array, all_products_db)

	//parse products struct to json format (accepted by dgraph)
	products_json, errorr := json.Marshal(new_prods)
	if errorr != nil {
		log.Fatal(errorr)
	}

	//creating a mutation transaction
	mu := &api.Mutation{
		CommitNow: true,
	}
	mu.SetJson = products_json
	//the assigned.Uids is a map[_:productId][uid] for the uploaded data
	assigned, e := db.NewTxn().Mutate(ctx, mu)
	if e != nil {
		log.Fatal(e)
	}
	//adding the new loaded products to the map that contains all of them
	for _, np := range new_prods {
		all_products_db[assigned.Uids[np.Uid]] = np.Uid
	}
	return all_products_db
}
