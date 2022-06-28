package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sophie-server/model"
)

// CreateUser Creates user instance and adds it to database
func CreateUser(c *gin.Context) {
	var userCreate model.UserCreate
	if err := c.BindJSON(&userCreate); err == nil {
		return
	}
	c.IndentedJSON(http.StatusCreated, userCreate)
}
