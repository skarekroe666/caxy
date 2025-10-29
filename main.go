package main

import (
	"caxy/cmd"
	"caxy/internal/server"
	"log"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalf("could not start the server: %v", err)
	}
	cmd.Execute()
}
