package main

import (
	"log"
	"net/http"
	"github.com/ortizdavid/golang-rest-api/config"
	"github.com/ortizdavid/golang-rest-api/controllers"
	"github.com/gorilla/mux"

)

func main() {

	router := mux.NewRouter()
	controllers.SetupRoutes(router)
	log.Println("Server Running on", config.ListenAddr())
	log.Fatal(http.ListenAndServe(config.ListenAddr(), router))
}