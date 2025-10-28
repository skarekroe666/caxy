package main

import (
	"caxy/internal/server"
	"log"
)

// var cunstomTransport = http.DefaultTransport

func main() {
	if err := server.Run(); err != nil {
		log.Fatalf("could not start the server: %v", err)
	}
}
