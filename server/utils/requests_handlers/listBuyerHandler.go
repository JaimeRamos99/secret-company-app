package request_handlers

import (
	json "encoding/json"
	log "log"
	http "net/http"

	logic "github.com/JaimeRamos99/prueba-truora-2/logic"
	structs "github.com/JaimeRamos99/prueba-truora-2/utils/structs"
	dgo "github.com/dgraph-io/dgo/v200"
)

func ListBuyersHandler(db *dgo.Dgraph, w http.ResponseWriter, r *http.Request) {
	struct_resp := logic.ListAllUsers(db)
	json_resp, err := json.Marshal(struct_resp)
	if err != nil {
		cli_resp_struct := *structs.NewClientResponse(true, "Something wrong happen")
		cli_resp_json, e := json.Marshal(cli_resp_struct)
		if e != nil {
			log.Fatal(e)
		}
		w.Write([]byte(cli_resp_json))
		log.Fatal(err)
	}
	w.Write([]byte(json_resp))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
