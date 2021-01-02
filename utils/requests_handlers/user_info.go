package request_handlers

import (
	http "net/http"

	"github.com/JaimeRamos99/prueba-truora-2/logic"
	dgo "github.com/dgraph-io/dgo/v200"
	chi "github.com/go-chi/chi"
)

func GetUserInfo(db *dgo.Dgraph, w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "userId")
	if id != "" {
		logic.UserInfo(db, id)
		w.Write([]byte("ok"))
	}
	w.Write([]byte("Invalid request"))
}
