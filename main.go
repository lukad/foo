package main

import (
	"os"

	"github.com/lukad/helix/web"

	"github.com/op/go-logging"
)

var (
	address string = "0.0.0.0"
	port    string = "8080"
	log     *logging.Logger
)

func init() {
	log = logging.MustGetLogger("main")

	if envAddress := os.Getenv("ADDRESS"); envAddress != "" {
		address = envAddress
	}
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
}

func main() {
	server := web.NewServer()
	if err := server.ListenAndServe(address + ":" + port); err != nil {
		log.Fatal(err)
	}
}
