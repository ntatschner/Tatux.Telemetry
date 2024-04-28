package main

import (
	"log"
	"net/http"
	"github.com/labstack/echo/v4"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

// Connection parameters

func connectToInfluxDB() {
	const (
		influxDBUrl := os.Getenv("INFLUXDB_URL")
		influxDBToken := os.Getenv("INFLUXDB_TOKEN")
		influxDBOrg := os.Getenv("INFLUXDB_ORG")
		influxDBBucket := os.Getenv("INFLUXDB_BUCKET")
		influxDBClient = influxdb2.NewClient(influxDBUrl, influxDBToken)
	)
	log.Println("Connected to InfluxDB")
}

// InfluxDB client
var influxDBClient influxdb2.Client

// Connect to InfluxDB
func connectToInfluxDB() {
	influxDBClient = influxdb2.NewClient(influxDBUrl, influxDBToken)
	log.Println("Connected to InfluxDB")
}

// Add data to InfluxDB
func addToInfluxDB() {
	// Get non-blocking write client
	writeAPI := influxDBClient.WriteAPIBlocking(influxDBOrg, influxDBBucket)
	// Create point
	p := influxdb2.NewPoint("stat",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg":24.5, "max":45},
	)
	// Write point immediately
	writeAPI.WritePoint(p)
	log.Println("Data added to InfluxDB")
}

// Handler for PUT method
func putHandler(c echo.Context) error {
	addToInfluxDB()
	return c.String(http.StatusOK, "Data added to InfluxDB")
}

func main() {
	// Connect to InfluxDB
	connectToInfluxDB()

	// Echo instance
	e := echo.New()

	// Routes
	e.PUT("/", putHandler)

	// Start server
	log.Println("Starting server on :1323")
	e.Start(":1323")
}