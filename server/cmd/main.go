package main

import (
	"log"

	"github.com/chomba-mumba/proxy-server-ad-block/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal("could nto start the server: %v", err)
	}
}
