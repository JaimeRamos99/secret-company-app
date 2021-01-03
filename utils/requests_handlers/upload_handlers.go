package request_handlers

import (
	json "encoding/json"
	http "net/http"
	regexp "regexp"
	strconv "strconv"
	time "time"

	logic "github.com/JaimeRamos99/prueba-truora-2/logic"
	utils "github.com/JaimeRamos99/prueba-truora-2/utils"
	dgo "github.com/dgraph-io/dgo/v200"
	redis "github.com/go-redis/redis/v8"
)

type InputDate struct {
	Date string `json:"date"`
}

func UploadHandler(db *dgo.Dgraph, rdb *redis.Client, w http.ResponseWriter, r *http.Request) {
	var input_date InputDate

	//Parse json to struct
	json.NewDecoder(r.Body).Decode(&input_date)
	match, _ := regexp.MatchString("^[0-9]{4}-[0-9]{2}-[0-9]{2}$", input_date.Date)

	//Checking if the Body param 'date' meets the expected format
	if match {
		//Parsing the input date string to the desired time format
		parsed_time, err := time.Parse(utils.LayoutISO, input_date.Date)
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

		//Valid date
		res := logic.UploadData(db, rdb, strconv.FormatInt(unix_time_input, 10))
		if res {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode("The data was upload")
			return
		}
		w.Write([]byte("The data had already been uploaded"))
		return
	}
	w.Write([]byte("Invalid date"))
}
