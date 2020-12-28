package logic

import (
  "fmt"
  "log"
  //"regexp"
  "strings"
  "net/http"
  "io/ioutil"
  "github.com/dgraph-io/dgo/v200"
)

func UploadTransactions(db *dgo.Dgraph, date string) bool{

  query_string := fmt.Sprintf(`https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/transactions?date=%v`, date)
  response, er := http.Get(query_string)
  if er != nil {
    log.Fatal(er)
    return false
  }

  responseData, err := ioutil.ReadAll(response.Body)
  if err != nil {
    log.Fatal(err)
    return false
  }

  data := string(responseData)
  rows := strings.Split(data, "#")
  fmt.Println(rows[1])
  return true
}
