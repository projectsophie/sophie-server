package controller

import (
	"github.com/gorilla/mux"
	routers "sophie-server/router"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	api := router.PathPrefix("/api").Subrouter()
	api = routers.InitUserRouter(api)
	return router
}
