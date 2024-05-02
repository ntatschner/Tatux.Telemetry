package app

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	ginrouter "github.com/gin-gonic/gin"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

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

// Handlers for GET methods

func getPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// Handler for PUT method

type Telemetry struct {
	Source        string    `json:"source"`
	Command       string    `json:"command"`
	Complete      bool      `json:"complete"`
	LocalDateTime time.Time `json:"localDateTime"`
	Exception     string    `json:"exception"`
}

func putTelemetry(c *gin.Context) {
	// Get the JSON from the request
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error reading request body",
		})
		return
	}

	// Unmarshal the JSON into the struct
	var telemetry Telemetry
	err = json.Unmarshal(body, &telemetry)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON",
			"cause": err.Error(),
		})
		return
	}
}

// Database Functions

func connectToInfluxDB(url string, token string) (*influxdb2.Client, error) {
	c := influxdb2.NewClient(url, token)
	return c
}

func Start() {
	// Starting the router
	router := ginrouter.Default()

	// Routes
	router.GET("/ping", getPing)
	router.PUT("/api/telemetry", putTelemetry)

	// Start server
	log.Println("Starting API server on :" + listenPort)
	router.Run(":" + listenPort)
}
