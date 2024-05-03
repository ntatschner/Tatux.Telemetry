package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetSystemHealth(c *gin.Context) {
	// Test internet connectivity
	var internet bool
	_, err := http.Get("https://www.google.com")
	if err != nil {
		internet = false
	} else {
		internet = true
	}
	// Test InfluxDB connectivity
	var influxdb bool
	_, err = http.Get(InfluxDBUrl)
	if err != nil {
		influxdb = false
	} else {
		influxdb = true
	}
	// Return the results as JSON
	c.JSON(http.StatusOK, gin.H{
		"internet": internet,
		"influxdb": influxdb,
	})
}

type Telemetry struct {
	Source        string    `json:"source"`
	Command       string    `json:"command"`
	Complete      bool      `json:"complete"`
	LocalDateTime time.Time `json:"localDateTime"`
	Exception     string    `json:"exception"`
}

func PutTelemetry(c *gin.Context) {
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

func GetPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}