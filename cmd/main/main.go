package main

import (
  "github.com/gorilla/mux"
  "net/http"
  "log"
  "github.com/ToffaKrtek/goCrud/pkg/routes"
  
)
func main(){
  r := mux.NewRouter()
  routes.RegisterStoreRoutes(r)
  http.Handle("/", r)
  log.Fatal(http.ListenAndServe("localhost:8080", r))
}
