package app

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	ginrouter "github.com/gin-gonic/gin"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

// InfluxDB connection parameters

var (
	influxDBUrl    = getEnv("INFLUXDB_URL", "", false) + ":" + getEnv("INFLUXDB_PORT", "", false)
	influxDBToken  = getEnv("INFLUXDB_TOKEN", "", false)
	influxDBOrg    = getEnv("INFLUXDB_ORG", "", false)
	influxDBBucket = getEnv("INFLUXDB_BUCKET", "", false)
	listenPort     = getEnv("LISTENONPORT", "9000", false)
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

// InfluxDB client
var influxDBClient influxdb2.Client

// Connection parameters

func connectToInfluxDB() {
	influxDBClient = influxdb2.NewClient(influxDBUrl, influxDBToken)
	log.Println("Connected to InfluxDB")
}

// Add data to InfluxDB

// Handlers for GET methods

func getPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func getHandler(c *gin.Context) error {
	return nil
}

// Handler for PUT method

func putHandler(c *gin.Context) error {
	return nil
}

func Start() {
	// Starting the router
	router := ginrouter.Default()

	// Routes
	router.GET("/ping", getPing)

	// Start server
	log.Println("Starting API server on :" + listenPort)
	router.Run(":" + listenPort)
}
