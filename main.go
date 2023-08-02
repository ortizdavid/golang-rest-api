package main

import (
	"log"
	"net/http"
	"github.com/ortizdavid/golang-rest-api/config"
	"github.com/ortizdavid/golang-rest-api/controllers"
)

func main() {

	controllers.SetupRoutes()
	log.Println("Server Running on", config.ListenAddr())
	http.ListenAndServe(config.ListenAddr(), nil)
}