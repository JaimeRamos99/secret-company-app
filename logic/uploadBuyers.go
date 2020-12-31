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

func UploadBuyers(db *dgo.Dgraph, date string) bool {

	url_string := fmt.Sprintf(utils.Base_url, "buyers", date)
	req, er := http.NewRequest("GET", url_string, nil)
	if er != nil {
		log.Fatal(er)
		return false
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer resp.Body.Close()

	var array_buyers []structs.Buyer
	error := json.NewDecoder(resp.Body).Decode(&array_buyers)
	if error != nil {
		log.Fatal(error)
		return false
	}
	return true
}
