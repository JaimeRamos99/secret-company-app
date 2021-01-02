package request_handlers

import (
	json "encoding/json"
	log "log"
	http "net/http"

	"github.com/JaimeRamos99/prueba-truora-2/logic"
	dgo "github.com/dgraph-io/dgo/v200"
	chi "github.com/go-chi/chi"
)

func GetUserInfo(db *dgo.Dgraph, w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "userId")
	if id != "" {
		all_info_struct := logic.UserInfo(db, id)

		resp_json, err := json.Marshal(all_info_struct)
		if err != nil {
			w.Write([]byte("Invalid request"))
			log.Fatal(err)
			return
		}
		w.Write([]byte(resp_json))
		return
	}
	w.Write([]byte("Invalid request"))
}
