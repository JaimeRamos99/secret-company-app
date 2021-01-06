package main

import (
	"fmt"
	log "log"
	http "net/http"
	time "time"

	cache "github.com/JaimeRamos99/prueba-truora-2/cache"
	database "github.com/JaimeRamos99/prueba-truora-2/database"
	logic "github.com/JaimeRamos99/prueba-truora-2/logic"
	utils "github.com/JaimeRamos99/prueba-truora-2/utils"
	request_handlers "github.com/JaimeRamos99/prueba-truora-2/utils/requests_handlers"
	chi "github.com/go-chi/chi"
	cors "github.com/go-chi/cors"
	middleware "github.com/go-chi/chi/middleware"
)

func init() {
	fmt.Println("-----------------------------------------------")
	//create dgraph schemas
	db, close_conn := database.NewClient()
	database.CreateEntitiesSchemas(db)

	//delete previous redis keys
	//cache.DeleteAllPreviousKeys()

	//upload today's data by default
	now := time.Now()
	date := now.Format(utils.LayoutISO)
	rdb := cache.NewClient()
	status := logic.UploadData(db, rdb, date)
	if status {
		log.Print("Today's data loaded succesfully")
	} else {
		log.Print("today's data had already been loaded")
	}
	defer close_conn()
}

func main() {

	//instance of the connection to the DB
	db, close_conn := database.NewClient()
	defer close_conn()

	//instance of redis cache
	rdb := cache.NewClient()

	// HTTP router and middleware initialization
	r := chi.NewRouter()
	 r.Use(cors.Handler(cors.Options{
    // AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
    AllowedOrigins:   []string{"*"},
    // AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300, // Maximum value not ignored by any of major browsers
  }))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	//endpoint to upload the transactions, buyers and products of a given day
	r.Post("/upload", func(w http.ResponseWriter, r *http.Request) {
		request_handlers.UploadHandler(db, rdb, w, r)
	})

	//endpoint to get all the users in the db
	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		request_handlers.ListBuyersHandler(db, w, r)
	})

	//endpoint that return all the info of a given user
	r.Route("/user_info", func(r chi.Router) {
		// Subrouter
		r.Route("/{userId}", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				request_handlers.GetUserInfo(db, rdb, w, r)
			})
		})
	})

	//service listening on port 3000
	http.ListenAndServe(":3000", r)
}
