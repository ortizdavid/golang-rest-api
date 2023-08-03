package controllers

import "github.com/gorilla/mux"

func SetupRoutes(router *mux.Router) {
	TaskController{}.RegisterRoutes(router)
}