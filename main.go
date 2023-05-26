package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/medium.rip/api"
)

func main() {
	log.Println("Starting HTTP server...")
	api.RegisterRoutes()
}
