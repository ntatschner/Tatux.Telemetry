package handlers

import (
	"encoding/json"
	"io"
	"log"
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
	ID                        string        `json:"id"`
	CommandName               string        `json:"commandname"`
	Complete                  bool          `json:"complete"`
	LocalDateTime             time.Time     `json:"localDateTime"`
	ExecutionDuration         time.Duration `json:"executionDuration"`
	Failed                    bool          `json:"failed"`
	Exception                 string        `json:"exception"`
	Manufacturer              string        `json:"Manufacturer"`
	Model                     string        `json:"Model"`
	TotalPhysicalMemory       int           `json:"TotalPhysicalMemory"`
	NumberOfProcessors        int           `json:"NumberOfProcessors"`
	NumberOfLogicalProcessors int           `json:"NumberOfLogicalProcessors"`
	PartOfDomain              bool          `json:"PartOfDomain"`
	HardwareSerialNumber      string        `json:"HardwareSerialNumber"`
	BootDriveSerial           string        `json:"BootDriveSerial"`
	OSType                    string        `json:"ostype"`
	OSArchitecture            string        `json:"OSArchitecture"`
	OSVersion                 string        `json:"Version"`
	OSBuildNumber             string        `json:"BuildNumber"`
	PowerShellVersion         string        `json:"PowerShellVersion"`
	HostVersion               string        `json:"HostVersion"`
	HostName                  string        `json:"HostName"`
	HostUI                    string        `json:"HostUI"`
	HostCulture               string        `json:"HostCulture"`
	HostUICulture             string        `json:"HostUICulture"`
	ModuleName                string        `json:"ModuleName"`
	ModuleVersion             string        `json:"ModuleVersion"`
	ModulePath                string        `json:"ModulePath"`
}

func GetPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
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
		log.Printf("Error: %v", err.Error())
		return
	} else {
		// Write the telemetry to InfluxDB
		WriteTelemetry(telemetry)
	}
}
