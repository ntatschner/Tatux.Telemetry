package main

import (
	"context"
	"log"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
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

var client = influxdb2.NewClient(influxDBUrl, influxDBToken)

func ensureConnection(url string, token string) {
	for {
		// Ping the InfluxDB server
		_, err := client.Health(context.Background())
		if err != nil {
			log.Println("Failed to ping InfluxDB:", err)

			// Attempt to reconnect
			log.Println("Attempting to reconnect to InfluxDB...")
			client = influxdb2.NewClient(url, token)
		}

		// Wait for a while before pinging again
		time.Sleep(30 * time.Second)
	}
}

func main() {
	log.Println("Starting server")

	client = influxdb2.NewClient(influxDBUrl, influxDBToken)

	go ensureConnection(influxDBUrl, influxDBToken)

	app.Start()
}
