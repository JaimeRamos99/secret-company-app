package request_handlers

import (
	json "encoding/json"
	log "log"
	http "net/http"

	logic "github.com/JaimeRamos99/prueba-truora-2/logic"
	dgo "github.com/dgraph-io/dgo/v200"
)

func ListBuyersHandler(db *dgo.Dgraph, w http.ResponseWriter, r *http.Request) {
	struct_resp := logic.ListAllUsers(db)
	json_resp, err := json.Marshal(struct_resp)
	if err != nil {
		w.Write([]byte("Something wrong happen"))
		log.Fatal(err)
	}
	w.Write([]byte(json_resp))
}
