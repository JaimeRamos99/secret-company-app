package logic

import (
  "fmt"
  "log"
  "encoding/json"
  "github.com/dgraph-io/dgo/v200"
  "github.com/JaimeRamos99/prueba-truora-2/database"
)

type Resp struct {
    Query []string `json:"query"`
}

func UploadData(db *dgo.Dgraph, date string) bool{
  res := database.CheckDate(db, date)
  res_string := fmt.Sprintf("%s\n", res.Json)

  var resp []Resp
	json.Unmarshal([]byte(res_string), &resp)
  log.Print(resp == nil)
  return true
}

func UploadBuyers(){

}

func UploadProducts(){

}

func UploadTransactions(){

}

func AddUploadDate(){

}
