package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

// Connection parameters
const (
	influxDBUrl = "http://localhost:8086"
	influxDBToken = "your-token"
	influxDBOrg = "your-org"
	influxDBBucket = "your-bucket"
)

// InfluxDB client
var influxDBClient influxdb2.Client

// Connect to InfluxDB
func connectToInfluxDB() {
	influxDBClient = influxdb2.NewClient(influxDBUrl, influxDBToken)
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
	e.Start(":1323")
}