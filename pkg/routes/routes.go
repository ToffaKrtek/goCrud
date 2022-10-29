package routes

import (
  "github.com/gorilla/mux"
  "github.com/ToffaKrtek/goCrud/pkg/controllers"
)
var RegisterStoreRoutes = func(router *mux.Router){
  router.HandleFunc("/items/", controllers.GetAll).Methods("GET")
  router.HandleFunc("/items/", controllers.Create).Methods("POST")
  router.HandleFunc("/items/", controllers.GetAll).Methods("GET")
  router.HandleFunc("/items/", controllers.GetAll).Methods("GET")
  router.HandleFunc("/items/", controllers.GetAll).Methods("GET")
}
