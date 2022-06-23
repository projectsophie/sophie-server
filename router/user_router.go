package router

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// CreateUser Creates user instance and adds it to database
func CreateUser(writer http.ResponseWriter, req *http.Request) {
	_ = json.NewEncoder(writer).Encode("уютненько")
}

// InitUserRouter Initializes a sub-router for router
// which was passed as an argument.
func InitUserRouter(router *mux.Router) {
	subRouter := router.PathPrefix("/users").Subrouter()
	subRouter.HandleFunc("/test", CreateUser)
}
