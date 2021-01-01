package logic

import (
	json "encoding/json"
	fmt "fmt"
	log "log"
	http "net/http"

	utils "github.com/JaimeRamos99/prueba-truora-2/utils"
	structs "github.com/JaimeRamos99/prueba-truora-2/utils/structs"
	dgo "github.com/dgraph-io/dgo/v200"
)

func UploadUsers(db *dgo.Dgraph, date string) bool {

	//creating the request to fetch the resource
	url_string := fmt.Sprintf(utils.Base_url, "buyers", date)
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

	//decode json into the equivalent struct
	var array_buyers []structs.Buyer
	error := json.NewDecoder(resp.Body).Decode(&array_buyers)
	if error != nil {
		log.Fatal(error)
		return false
	}
	all_users_db := GetAllUsers(db)
	fmt.Println(all_users_db)
	return true
}
