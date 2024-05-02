package main

import (
	"log"
	"os"

	"github.com/ntatschner/Tatux.Telemetry/src/api/app"
)

var (
	influxDBUrl   = getEnv("INFLUXDB_URL", "", false) + ":" + getEnv("INFLUXDB_PORT", "", false)
	influxDBToken = getEnv("INFLUXDB_TOKEN", "", false)
)

func getEnv(key, defaultValue string, throwOnDefault bool) string {
	value, exists := os.LookupEnv(key)
	if !exists && !throwOnDefault {
		return defaultValue
	} else if !exists && throwOnDefault {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return value
}

func main() {
	log.Println("Starting server")

	go handlers.connectInfluxDB(influxDBUrl, influxDBToken)

	app.Start()
}
