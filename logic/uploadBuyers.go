package logic

import (
  "fmt"
  "log"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "github.com/dgraph-io/dgo/v200"
)

type Buyer struct {
  Id   string `json:"id"`
  Name string `json:"name"`
  Age  int    `json:"age"`
}
func UploadBuyers(db *dgo.Dgraph, date string) bool{
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
	var buyers []Buyer
	json.Unmarshal([]byte(data_str), &buyers)
  for _, reg := range buyers {
    log.Print(len(reg.Id))
  }

  return true
}
