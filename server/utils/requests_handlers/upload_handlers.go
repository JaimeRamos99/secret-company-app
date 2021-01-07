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
	match, _ := regexp.MatchString("^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])$", input_date.Date)

	//Checking if the Body param 'date' meets the expected format
	if match {
		//Parsing the input date string to the desired time format
		parsed_time, err := time.Parse(utils.LayoutISO, input_date.Date)
		if err != nil {
			response := utils.CreateResponse(true, "Invalid date")
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(response))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		//Getting the timestamp of the input date and validating if it's coherent
		unix_time_input := parsed_time.Unix()
		timestamp_now := time.Now().Unix()
		if unix_time_input < 0 || (unix_time_input > timestamp_now) {
			response := utils.CreateResponse(true, "Invalid date, it's the future")
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(response))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		//Valid date
		res := logic.UploadData(db, rdb, strconv.FormatInt(unix_time_input, 10))
		if res {
			response := utils.CreateResponse(false, "the data has been successfully uploaded")

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(response))
			return
		}
		response := utils.CreateResponse(false, "The data had already been uploaded")
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(response))
		w.WriteHeader(http.StatusOK)
		return
	}
	response := utils.CreateResponse(true, "Invalid date")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response))
	w.WriteHeader(http.StatusBadRequest)
}
