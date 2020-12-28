package logic

import (
  "fmt"
  "encoding/json"
  "github.com/dgraph-io/dgo/v200"
  "github.com/JaimeRamos99/prueba-truora-2/database"
)

//Struct for the query response
type Resp struct {
    Date string `json:"date"`
}

type RespArray struct {
  Query []Resp `json:"query"`
}

//Main function that handles the whole process of uploading the data of a given day
func UploadData(db *dgo.Dgraph, date string) bool{
  res := database.CheckDate(db, date)
  res_string := fmt.Sprintf("%s\n", res.Json)

  var resp RespArray
	json.Unmarshal([]byte(res_string), &resp)


  //The data has to be uploaded, because hasn't been uploaded
  //UploadProducts(db, date)
  //UploadBuyers(db, date)
  UploadTransactions(db, date)
  /*if len(resp.Query) == 0 {
    UploadTransactions(db, date)
    database.AddUploadDate(db, date)
    return true
    //or false
  }*/

  return false
}
