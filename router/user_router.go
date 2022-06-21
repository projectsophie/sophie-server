package router

import (
	"github.com/gorilla/mux"
)

// Router handler

func HandleRouter() {
	router := mux.NewRouter().StrictSlash(true)
	api := router.PathPrefix("/api").Subrouter()

}
