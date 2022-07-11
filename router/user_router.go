package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sophie-server/middleware/session"
	mod "sophie-server/model/pages/types"
	"sophie-server/service"
	"sophie-server/store"
)

// CreateUser Creates users instance and adds it to database
func CreateUser(c *gin.Context) {
	//var userCreate users.UserCreate
	//if err := c.BindJSON(&userCreate); err != nil {
	//	return
	//}
	//c.IndentedJSON(http.StatusOK, store.CreateUser(userCreate))
	table := mod.Table{ID: 1, Title: "Test", Workspace: 1, Type: "Table", Columns: []map[string]interface{}{{"dd": "dddd", "ddf": 2}}, Strings: []map[string]interface{}{{"Test": 2, "Test2": 3}}}
	test := table.GetMetadata()
	c.JSON(http.StatusOK, test)
}

// AuthUser is a method that generates a JWT token for users via gin's context
// or rejects the request if users data is not valid and not found in database.
func AuthUser(c *gin.Context) {
	token := service.Login(c)
	if token != "" {
		store.AppendSession(session.GenerateSession(c, token))
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, nil)
	}
}

// GetCurrentUser is a method that returns
// current users instance via gin's context.
func GetCurrentUser(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	tokenRaw := authHeader[len("Bearer "):]
	userGet, success := store.GetAuthedUser(tokenRaw)
	if success {
		c.IndentedJSON(http.StatusOK, userGet)
	} else {
		c.JSON(http.StatusUnauthorized, nil)
	}
}

// Logout is a method that logs out users via gin's context.
func Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	tokenRaw := authHeader[len("Bearer "):]
	store.DeleteSession(tokenRaw)
	c.JSON(http.StatusOK, nil)
}
