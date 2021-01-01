package logic

import (
	json "encoding/json"
	fmt "fmt"

	database "github.com/JaimeRamos99/prueba-truora-2/database"
	structs "github.com/JaimeRamos99/prueba-truora-2/utils/structs"
	dgo "github.com/dgraph-io/dgo/v200"
)

//Main function that handles the whole process of uploading the data of a given day
func UploadData(db *dgo.Dgraph, date string) bool {
	res := database.CheckDate(db, date)
	res_string := fmt.Sprintf("%s\n", res.Json)

	var resp *structs.RespArray
	json.Unmarshal([]byte(res_string), &resp)

	//The data has to be uploaded, because hasn't been uploaded

	//if len(resp.Query) == 0 {
	usr_map := UploadUsers(db, date)
	prods_map := UploadProducts(db, date)
	HandleTransactions(db, date, usr_map, prods_map)

	//database.AddUploadDate(db, date)
	//return true
	//or false
	//}

	return /*false*/ true
}
