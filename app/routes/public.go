package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// SetPublicRouter expone los archivos publicos
func SetPublicRouter(router *mux.Router) {
	router.Handle("/", http.FileServer(http.Dir("./public")))
}
