package routes

import (
	"github.com/gorilla/mux"
)

// InitRoutes inicia todas las rutas
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	SetloginRouter(router)
	SetUserRouter(router)
	SetCommentRouter(router)
	SetVoteRouter(router)
	SetRealTimeRouter(router)
	SetPublicRouter(router)

	return router
}
