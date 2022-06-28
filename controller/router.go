package controller

import (
	"github.com/gin-gonic/gin"
	"sophie-server/store"
)

var router *gin.Engine

func GetRouter() *gin.Engine {
	if router != nil {
		return router
	}
	panic("Router not initialized. Seems, that InitRouter() hasn't been executed or router hasn't been initialized.")
}

func InitUserRoutes() {
	router.GET("/api/test", store.Test)
	router.POST("/api/users/create", store.CreateUser)
}

func InitRouter() {
	router = gin.Default()
	InitUserRoutes()
}
