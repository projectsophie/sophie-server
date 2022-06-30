package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sophie-server/model"
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
