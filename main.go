package main

import (
	"log"
	"net/http"
	"sophie-server/controller"
	"sophie-server/database"
	"strconv"
)

func main() {
	log.Println("Initializing a Sophie server...")
	log.Println("Initializing databases...")
	database.InitDB(true, true)
	log.Println("Databases were initialized successfully.")
	log.Println("Initializing server...")
	InitServer("127.0.0.1", 8080)
}

// InitServer Initializes server and routers
func InitServer(host string, port int) {
	address := host + ":" + strconv.Itoa(port)
	log.Printf("Sophie server listening on http://%s", address)
	err := http.ListenAndServe(address, controller.InitRouter())
	if err != nil {
		panic("Couldn't initialize server due error.")
	}
}
