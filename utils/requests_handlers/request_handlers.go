package request_handlers

import (
  "time"
  "regexp"
  "net/http"
  "encoding/json"
  "github.com/dgraph-io/dgo/v200"
  "github.com/JaimeRamos99/prueba-truora-2/logic"
  //"github.com/JaimeRamos99/prueba-truora-2/database"
)

type InputDate struct {
  Date   string `json: "date"`
}

const (
	layoutISO = "2006-01-02"
)

func UploadHandler(db *dgo.Dgraph, w http.ResponseWriter, r *http.Request){
  var input_date InputDate
  json.NewDecoder(r.Body).Decode(&input_date)
  match, _ := regexp.MatchString("^[0-9]{4}-[0-9]{2}-[0-9]{2}$", input_date.Date)

  //Checking if the Body param 'date' meets the expected format
  if match {
    //Parsing the input date string to the desired time format
    parsed_time, err := time.Parse(layoutISO, input_date.Date)
    if err != nil {
        w.Write([]byte("Invalid date"))
        return
    }

    //Getting the timestamp of the input date and validating if it's coherent
    unix_time_input := parsed_time.Unix()
    timestamp_now := time.Now().Unix()
    if unix_time_input < 0 || (unix_time_input > timestamp_now) {
      w.Write([]byte("Invalid date"))
      return
    }
    res := logic.UploadData(db, input_date.Date)
    if res {
      w.Write([]byte("The data was upload"))
      return
    }
    w.Write([]byte("The data had already been uploaded"))
    return
  }
  w.Write([]byte("Invalid date"))
}
