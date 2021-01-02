package main

import (
	"log"
	http "net/http"
	time "time"

	database "github.com/JaimeRamos99/prueba-truora-2/database"
	logic "github.com/JaimeRamos99/prueba-truora-2/logic"
	utils "github.com/JaimeRamos99/prueba-truora-2/utils"
	request_handlers "github.com/JaimeRamos99/prueba-truora-2/utils/requests_handlers"
	chi "github.com/go-chi/chi"
	middleware "github.com/go-chi/chi/middleware"
)

func init() {
	//create dgraph schemas
	db, close_conn := database.NewClient()
	database.CreateEntitiesSchemas(db)
	now := time.Now()
	date := now.Format(utils.LayoutISO)
	logic.UploadData(db, date)
	log.Print("Today's data loaded succesfully")
	defer close_conn()
}

func main() {

	//instance of the connection to the DB
	db, close_conn := database.NewClient()
	defer close_conn()

	//HTTP router and middleware initialization
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	//Endpoint to upload the transactions, buyers and products of a given day
	r.Post("/upload", func(w http.ResponseWriter, r *http.Request) {
		request_handlers.UploadHandler(db, w, r)
	})

	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		request_handlers.ListBuyersHandler(db, w, r)
	})

	r.Route("/user_info", func(r chi.Router) {

		// Subrouter
		r.Route("/{userId}", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				request_handlers.GetUserInfo(db, w, r)
			})
		})
	})

	//service listening on port 3000
	http.ListenAndServe(":3000", r)
}
