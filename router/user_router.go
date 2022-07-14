package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sophie-server/middleware/auth"
	"sophie-server/middleware/session"
	"sophie-server/model/users"
	"sophie-server/store"
	"sophie-server/util"
)

// CreateUser Creates users instance and adds it to database
func CreateUser(c *gin.Context) {
	var userCreate users.UserCreate
	if err := c.BindJSON(&userCreate); err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, store.CreateUser(userCreate))
}

// AuthUser is a method that generates a JWT token for users via gin's context
// or rejects the request if users data is not valid and not found in database.
func AuthUser(c *gin.Context) {
	token := auth.Login(c)
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
	userGet, success := store.GetAuthedUser(util.GetToken(c))
	if success {
		c.IndentedJSON(http.StatusOK, userGet)
	} else {
		c.JSON(http.StatusUnauthorized, nil)
	}
}

// Logout is a method that logs out users via gin's context.
func Logout(c *gin.Context) {
	store.DeleteSession(util.GetToken(c))
	c.JSON(http.StatusOK, nil)
}

// GetUserWorkspaces is a method that returns
// current users workspaces via gin's context.
func GetUserWorkspaces(c *gin.Context) {
	userGet, success := store.GetAuthedUser(util.GetToken(c))
	if success {
		c.IndentedJSON(http.StatusOK, userGet.Workspaces)
	} else {
		c.JSON(http.StatusUnauthorized, nil)
	}
}
