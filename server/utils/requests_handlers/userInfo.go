package request_handlers

import (
	json "encoding/json"
	log "log"
	http "net/http"

	logic "github.com/JaimeRamos99/prueba-truora-2/logic"
	utils "github.com/JaimeRamos99/prueba-truora-2/utils"
	dgo "github.com/dgraph-io/dgo/v200"
	chi "github.com/go-chi/chi"
	redis "github.com/go-redis/redis/v8"
)

//IMPORT UTILS
func GetUserInfo(db *dgo.Dgraph, rdb *redis.Client, w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "userId")
	resp_invalid := utils.CreateResponse(true, "Invalid date")
	if id != "" {
		all_info_struct := logic.UserInfo(db, rdb, id)
		resp_json, err := json.Marshal(all_info_struct)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(resp_invalid))
			w.WriteHeader(http.StatusBadRequest)
			log.Fatal(err)
			return
		}

		w.Write([]byte(resp_json))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(resp_invalid))
	w.WriteHeader(http.StatusBadRequest)
}
