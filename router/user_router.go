package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"sophie-server/database"
)

// CreateUser Creates user instance and adds it to database
func CreateUser(writer http.ResponseWriter, req *http.Request) {
	statement, _ := database.GetUsersDB().Prepare(database.CREATE_USER)
	_, _ = statement.Exec("Maxim", "Bataron", "dion", "test123", "test@gmail.com")
}

// InitUserRouter Initializes a sub-router for router
// which was passed as an argument.
func InitUserRouter(router *mux.Router) {
	subRouter := router.PathPrefix("/users").Subrouter()
	subRouter.HandleFunc("/test", CreateUser)
}
