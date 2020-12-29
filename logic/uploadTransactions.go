package logic

import (
  "fmt"
  "log"
  "strings"
  "unicode"
  "net/http"
  "io/ioutil"
  "github.com/dgraph-io/dgo/v200"
  "github.com/JaimeRamos99/prueba-truora-2/utils"
)

//Given a transaction string, this func converts it
//into a valid Transaction struct
func splitTransactions(tr string) Transaction{
  allowed_runes := utils.LoopDigits()
  acum := ""

  //creating a mirror string, except for the especial rune
  //instead a _ is added, in order to split later
  for _, rune := range tr {
    _, is_digit_or_special := allowed_runes[string(rune)]
    if (unicode.IsLetter(rune) || is_digit_or_special) {
      acum = acum + string(rune)
    }else{
      acum = acum + "_"
    }
  }

  //splitting transactions attrs
  tran_splitted := strings.Split(acum, "_")
  len_prods_str := len(tran_splitted[4])
  products  := strings.Split(tran_splitted[4][1:len_prods_str-1], ",")

  //Creating the Products array struct
  var products_Array []Product
  for _, product := range products {
    prod := Product{
      ProductId: product,
    }
    products_Array = append(products_Array, prod)
  }

  //instance of Transaction
  tran := Transaction{
    IdTransaction: tran_splitted[0],
    BuyerId: tran_splitted[1],
    Ip: tran_splitted[2],
    Device: tran_splitted[3],
    Products: products_Array,
  }

  return tran
}

//Function that handles the process of insert a transaction into the db
func UploadTransactions(db *dgo.Dgraph, date string) bool{
  //create the request
  url_string := fmt.Sprintf(`https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/[1]v?date=%[2]v`, "transactions", date)
  req, er := http.NewRequest("GET", url_string, nil)
  if er != nil {
    log.Fatal(er)
    return false
  }

  //commit the request
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    log.Fatal(err)
    return false
  }
  defer resp.Body.Close()

  //read the body response (our transactions data) and store as bytes
  bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

  //parse de data in bytes format to string
  //and split each transaction info by #
  input_str := string(bytes)
  transactionss := strings.Split(input_str,"#")
  transactions := transactionss[1:]

  //Sending every transaction string, to get a valid transaction struct
  var trans []Transaction
  for _, tr := range transactions {
    tran := splitTransactions(tr)
    trans = append(trans, tran)
  }
  return true
}
