package main

import (
  "net/http"
  "github.com/go-chi/chi"
  "github.com/go-chi/chi/middleware"
  "github.com/JaimeRamos99/prueba-truora-2/database"
  "github.com/JaimeRamos99/prueba-truora-2/utils/requests_handlers"
)

func init(){
  //create dgraph schemas
  db, close_conn := database.NewClient()
  database.CreateSchemaUploadedDates(db)
  database.CreateEntitiesSchemas(db)
  defer close_conn()
}

func main(){

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
  r.Post("/upload", func(w http.ResponseWriter, r *http.Request){
    request_handlers.UploadHandler(db, w, r)
  })

  //service listening on port 3000
  http.ListenAndServe(":3000", r)
}
