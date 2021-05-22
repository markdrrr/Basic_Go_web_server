package main

import (
	"log"
	"server_go/server"
)

func main() {
	newServer := server.NewServer("localhost", 8001)
	err := newServer.Run()
	if err != nil {
		log.Fatal(err)
	}
}
