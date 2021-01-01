package logic

import (
	context "context"
	json "encoding/json"
	fmt "fmt"
	log "log"
	http "net/http"

	utils "github.com/JaimeRamos99/prueba-truora-2/utils"
	structs "github.com/JaimeRamos99/prueba-truora-2/utils/structs"
	dgo "github.com/dgraph-io/dgo/v200"
	api "github.com/dgraph-io/dgo/v200/protos/api"
)

func UploadUsers(db *dgo.Dgraph, date string) map[string]string {

	ctx := context.Background()

	//creating the request to fetch the resource
	url_string := fmt.Sprintf(utils.Base_url, "buyers", date)
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

	//decode json into the equivalent struct
	var array_buyers []structs.Buyer
	error := json.NewDecoder(resp.Body).Decode(&array_buyers)
	if error != nil {
		log.Fatal(error)
	}

	//parse input obj format to a better to handle one
	var array_users structs.Users
	for _, byr := range array_buyers {
		usr := *structs.NewUser(byr.Id, byr.Name, byr.Age)
		array_users.Users = append(array_users.Users, usr)
	}
	all_users_db := GetAllUsers(db)
	new_users := NewUsers(array_users, all_users_db)

	//parse Users struct to json format (accepted by dgraph)
	users_json, errorr := json.Marshal(new_users.Users)
	if errorr != nil {
		log.Fatal(errorr)
	}

	//mutation object for dgo
	mu := &api.Mutation{
		CommitNow: true,
	}
	mu.SetJson = users_json

	//the assigned.Uids is a map[_:productId][uid] for the uploaded data
	assigned, err := db.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	//adding the new loaded products to the map that contains all of them
	for _, nu := range new_users.Users {
		all_users_db[assigned.Uids[nu.Uid]] = nu.Uid
	}
	return all_users_db
}
