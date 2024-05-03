package main

import (
	"log"

	"github.com/ntatschner/Tatux.Telemetry/src/api/app"
	"github.com/ntatschner/Tatux.Telemetry/src/api/handlers"
	"github.com/ntatschner/Tatux.Telemetry/src/api/system"
)

var (
	InfluxDBUrl   = system.GetEnv("INFLUXDB_URL", "", false) + ":" + system.GetEnv("INFLUXDB_PORT", "", false)
	InfluxDBToken = system.GetEnv("INFLUXDB_TOKEN", "", false)
)

func main() {
	log.Println("Starting server")

	go handlers.ConnectInfluxDB(InfluxDBUrl, InfluxDBToken)

	app.Start()
}
