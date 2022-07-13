package controller

import (
	"github.com/gin-gonic/gin"
	"sophie-server/middleware/auth"
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

// InitUserRoutes initializes users routes.
func InitUserRoutes() {
	userRoutes := router.Group("/api/users", auth.AuthorizeJWT())
	{
		userRoutes.GET("/me", routers.GetCurrentUser)
		userRoutes.GET("/logout", routers.Logout)
		userRoutes.GET("/workspaces", routers.GetUserWorkspaces)
	}
	router.GET("/api/users/create", routers.CreateUser)
	router.POST("/api/users/auth", routers.AuthUser)
}

// InitRouter initializes main router.
func InitRouter() {
	router = gin.Default()
	InitUserRoutes()
}
