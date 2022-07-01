package controller

import (
	"github.com/gin-gonic/gin"
	routers "sophie-server/router"
)

// router is a main router.
var router *gin.Engine

// GetRouter returns main router or closes server with error
// if router hasn't been initialized.
func GetRouter() *gin.Engine {
	if router != nil {
		return router
	}
	panic("Router not initialized. Seems, that InitRouter() hasn't been executed or router hasn't been initialized.")
}

// InitUserRoutes initializes user routes.
func InitUserRoutes() {
	router.POST("/api/users/create", routers.CreateUser)
	router.POST("/api/users/auth", routers.AuthUser)
}

// InitRouter initializes main router.
func InitRouter() {
	router = gin.Default()
	InitUserRoutes()
}
