package logic

import (
	context "context"
	fmt "fmt"
	ioutil "io/ioutil"
	log "log"
	http "net/http"
	strings "strings"

	utils "github.com/JaimeRamos99/prueba-truora-2/utils"
	structs "github.com/JaimeRamos99/prueba-truora-2/utils/structs"
	dgo "github.com/dgraph-io/dgo/v200"
	api "github.com/dgraph-io/dgo/v200/protos/api"
)

//Function that handles the process of insert a transaction into the db
func HandleTransactions(db *dgo.Dgraph, date string, usrs_map map[string]string, prods_map map[string]string) bool {

	ctx := context.Background()
	//create the request, passing the /name and the date
	url_string := fmt.Sprintf(utils.Base_url, "transactions", date)
	req, er := http.NewRequest("GET", url_string, nil)
	if er != nil {
		log.Fatal(er)
		return false
	}

	//commit the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer resp.Body.Close()

	//read the body response (our transactions data) and store as bytes
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return false
	}

	//parse de data in bytes format to string
	//and split each transaction info by #
	input_str := string(bytes)
	transactionss := strings.Split(input_str, "#")
	transactions := transactionss[1:]

	//Sending every transaction string, to get a valid transaction struct
	var trans []structs.Transaction
	for _, tr := range transactions {
		trans = append(trans, splitTransactions(tr))
	}

	transactions_json := CreateTransactionJson(trans, usrs_map, prods_map)

	//creating a mutation transaction
	mu := &api.Mutation{
		CommitNow: true,
	}
	mu.SetJson = transactions_json

	_, e := db.NewTxn().Mutate(ctx, mu)
	if e != nil {
		log.Fatal(e)
	}
	return true
}
