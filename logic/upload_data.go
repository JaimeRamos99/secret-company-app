package logic

import (
  "fmt"
  "encoding/json"
  "github.com/dgraph-io/dgo/v200"
  "github.com/JaimeRamos99/prueba-truora-2/database"
)

//Struct for the query response
type Resp struct {
    Query []string `json:"query"`
}

//Main function that handles the whole process of uploading the data of a given day
func UploadData(db *dgo.Dgraph, date string) bool{
  res := database.CheckDate(db, date)
  res_string := fmt.Sprintf("%s\n", res.Json)
  fmt.Println(res_string)
  var resp []Resp
	json.Unmarshal([]byte(res_string), &resp)
  //The data has to be uploaded
  if resp == nil {
    database.AddUploadDate(db, date)
    return true
    //or false
  }

  return false
}

func UploadBuyers(){

}

func UploadProducts(){

}

func UploadTransactions(){

}
