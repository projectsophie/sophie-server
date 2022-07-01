package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sophie-server/model"
	"sophie-server/service"
	"sophie-server/store"
)

// CreateUser Creates user instance and adds it to database
func CreateUser(c *gin.Context) {
	var userCreate model.UserCreate
	if err := c.BindJSON(&userCreate); err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, store.CreateUser(userCreate))
}

// AuthUser is a method that generates a JWT token for user via gin's context
// or rejects the request if user data is not valid and not found in database.
func AuthUser(c *gin.Context) {
	token := service.Login(c)
	if token != "" {
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, nil)
	}
}
