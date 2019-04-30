package routes

import (
	"../controllers"
	"github.com/gorilla/mux"
)

// SetloginRouter router para login
func SetloginRouter(router *mux.Router) {
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
}
