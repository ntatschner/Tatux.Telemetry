package app

import (
	"log"
	"net/http"
	"os"

	ginrouter "github.com/gin-gonic/gin"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/labstack/echo/v4"
)

// InfluxDB connection parameters

var (
	influxDBUrl    = os.Getenv("INFLUXDB_URL") + ":" + os.Getenv("INFLUXDB_PORT")
	influxDBToken  = os.Getenv("INFLUXDB_TOKEN")
	influxDBOrg    = os.Getenv("INFLUXDB_ORG")
	influxDBBucket = os.Getenv("INFLUXDB_BUCKET")
	listenPort     = os.Getenv("LISTENONPORT")
)

// InfluxDB client
var influxDBClient influxdb2.Client

// Connection parameters

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
		map[string]interface{}{"avg": 24.5, "max": 45},
	)
	// Write point immediately
	writeAPI.WritePoint(p)
	log.Println("Data added to InfluxDB")
}

// Handler for GET method

func getHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// Handler for PUT method
func putHandler(c echo.Context) error {
	addToInfluxDB()
	return c.String(http.StatusOK, "Data added to InfluxDB")
}

func Start() {
	// Starting the router
	router := ginrouter.Default()

	// Routes
	router.PUT("/api", putHandler)
	router.GET("/api", getHandler)

	// Start server
	log.Println("Starting API server on :" + listenPort)
	router.Start(":" + listenPort)
}
