package main

import (
	"log"
	"sophie-server/controller"
	"sophie-server/database"
	"strconv"
)

const (
	host = "127.0.0.1" // host is an IP of Sophie server
	port = 8080        // port is a port of Sophie server
)

func main() {
	log.Println("Initializing a Sophie server...")
	log.Println("Initializing databases...")
	database.InitDB(true, true)
	log.Println("Databases were initialized successfully.")
	log.Println("Initializing server...")
	InitServer(host, port)
}

// InitServer Initializes server and routers
func InitServer(host string, port int) {
	address := host + ":" + strconv.Itoa(port)
	log.Printf("Sophie server listening on http://%s", address)
	controller.InitRouter()
	err := controller.GetRouter().Run(address)
	if err != nil {
		panic("Couldn't initialize server due error.")
	}
}
