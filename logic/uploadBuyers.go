package logic

import (
	json "encoding/json"
	fmt "fmt"
	ioutil "io/ioutil"
	log "log"
	http "net/http"

	structs "github.com/JaimeRamos99/prueba-truora-2/utils/structs"
	dgo "github.com/dgraph-io/dgo/v200"
)

func UploadBuyers(db *dgo.Dgraph, date string) bool {
	query_string := fmt.Sprintf(`https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers?date=%v`, date)
	response, er := http.Get(query_string)

	if er != nil {
		log.Fatal(er)
		return false
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return false
	}

	data_str := string(data)
	var buyers []structs.Buyer
	json.Unmarshal([]byte(data_str), &buyers)
	for _, reg := range buyers {
		log.Print(len(reg.Id))
	}

	return true
}
