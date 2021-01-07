package utils

import (
	json "encoding/json"
	log "log"

	structs "github.com/JaimeRamos99/prueba-truora-2/utils/structs"
)

func CreateResponse(err bool, message string) []byte {
	cli_resp_struct := *structs.NewClientResponse(err, message)
	cli_resp_json, e := json.Marshal(cli_resp_struct)
	if e != nil {
		log.Fatal(e)
	}
	return cli_resp_json
}
