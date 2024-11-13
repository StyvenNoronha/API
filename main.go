package main

import (
	"api/api"
	"log"
)


func main() {
	server:= api.NewServer()
	server.ConfigureRoutes()
	if err:= server.Start(); err != nil{
		log.Fatal(err)
	}
}

