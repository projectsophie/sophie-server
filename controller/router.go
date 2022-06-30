package controller

import (
	"github.com/gin-gonic/gin"
	controller "sophie-server/router"
)

var router *gin.Engine

func GetRouter() *gin.Engine {
	if router != nil {
		return router
	}
	panic("Router not initialized. Seems, that InitRouter() hasn't been executed or router hasn't been initialized.")
}

func InitUserRoutes() {
	router.POST("/api/users/create", controller.CreateUser)
}

func InitRouter() {
	router = gin.Default()
	InitUserRoutes()
}
