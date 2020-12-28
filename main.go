package main

import (
  "net/http"
  "github.com/go-chi/chi"
  "github.com/go-chi/chi/middleware"
  "github.com/JaimeRamos99/prueba-truora-2/database"
  "github.com/JaimeRamos99/prueba-truora-2/utils/requests_handlers"
)

func init(){
  db, close_conn := database.NewClient()
  database.CreateSchemaUploadedDates(db)
  defer close_conn()
}

func main(){
  db, close_conn := database.NewClient()
  defer close_conn()

  r := chi.NewRouter()
  r.Use(middleware.RequestID)
  r.Use(middleware.RealIP)
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)

  r.Post("/upload", func(w http.ResponseWriter, r *http.Request){
    request_handlers.UploadHandler(db, w, r)
  })
  http.ListenAndServe(":3000", r)
}
