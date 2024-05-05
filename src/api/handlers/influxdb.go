package handlers

import (
	"context"
	"log"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/ntatschner/Tatux.Telemetry/src/api/system"
)

var (
	InfluxDBUrl     = system.GetEnv("INFLUXDB_URL", "", false) + ":" + system.GetEnv("INFLUXDB_PORT", "", false)
	InfluxDBToken   = system.GetEnv("INFLUXDB_TOKEN", "", false)
	INFLUXDB_BUCKET = system.GetEnv("INFLUXDB_BUCKET", "", false)
	INFLUXDB_ORG    = system.GetEnv("INFLUXDB_ORG", "", false)
)

var client = influxdb2.NewClient(InfluxDBUrl, InfluxDBToken)

func ConnectInfluxDB(url string, token string) {
	for {
		// Ping the InfluxDB server
		_, err := client.Health(context.Background())
		if err != nil {
			log.Println("Failed to ping InfluxDB:", err)

			// Attempt to reconnect
			log.Println("Attempting to reconnect to InfluxDB...")
			client = influxdb2.NewClient(url, token)
			_, err = client.Health(context.Background())
			if err != nil {
				log.Println("Failed to reconnect to InfluxDB:", err)
			} else {
				log.Println("Reconnected to InfluxDB")
			}
		} else {
			//log.Println("InfluxDB is healthy")
		}

		// Wait for a while before pinging again
		time.Sleep(30 * time.Second)
	}
}

func WriteTelemetry(telemetry Telemetry) {
	// Create a new point
	point := influxdb2.NewPointWithMeasurement(telemetry.ID).
		AddTag("commandName", telemetry.CommandName).
		AddField("complete", telemetry.Complete).
		AddField("Failed", telemetry.Failed).
		AddField("localDateTime", telemetry.LocalDateTime).
		AddField("exception", telemetry.Exception).
		SetTime(time.Now())

	// Write the point
	writeAPI := client.WriteAPIBlocking(INFLUXDB_ORG, INFLUXDB_BUCKET)
	writeAPI.WritePoint(context.Background(), point)
}
