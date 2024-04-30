package main

import (
	"log"

	"github.com/ntatschner/Tatux.Telemetry/src/api/app"
)

func main() {
	log.Println("Starting server")
	app.Start()
}
