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

type TelemetryBasic struct {
	Source        string    `json:"source"`
	Command       string    `json:"command"`
	Complete      bool      `json:"complete"`
	LocalDateTime time.Time `json:"localDateTime"`
 ExecutionTime time.Duration `json:"executionTime"`
	Exception     string    `json:"exception"`
}

type TelemetryFull struct {
  Source        string    `json:"source"`
  Command       string    `json:"command"`
  Complete      bool      `json:"complete"`
  LocalDateTime time.Time `json:"localDateTime"`
  ExecutionTime time.Duration `json:"executionTime"`
  Exception     string    `json:"exception"`
  Manufacturer string `json:"Manufacturer"`
  Model string `json:"Model"`
  TotalPhysicalMemory string `json:"TotalPhysicalMemory"`
  NumberOfProcessors string `json:"NumberOfProcessors"`
  NumberOfLogicalProcessors string `json:"NumberOfLogicalProcessors"`
  OSArchitecture string `json:"OSArchitecture"`
  OSVersion string `json:"Version"`
  OSBuildNumber string `json:"BuildNumber"`
  ServicePackMajorVersion string `json:"ServicePackMajorVersion"`
  ServicePackMinorVersion string `json:"ServicePackMinorVersion"`
  PowerShellVersion string `json:"PowerShellVersion"`
  HostVersion string `json:"HostVersion"`
  HostName string `json:"HostName"`
  HostUI string `json:"HostUI"`
  HostCulture string `json:"HostCulture"`
  HostUICulture string `json:"HostUICulture"`
  ModuleName string `json:"ModuleName"`
  ModuleVersion string `json:"ModuleVersion"`
  ModulePath string `json:"ModulePath"`
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
	} else {
		// Write the telemetry to InfluxDB
		WriteTelemetry(telemetry)
	}
}

func GetPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

type GenericTelemetry struct {
    Type string          `json:"type"`
    Data json.RawMessage `json:"data"`
}

func PutTelemetry(c *gin.Context) {
// Unmarshal the JSON into the generic struct
var generic GenericTelemetry
err = json.Unmarshal(body, &generic)
if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
        "error": "Invalid JSON",
        "cause": err.Error(),
    })
    return
}

// Determine which struct to unmarshal the data into based on the Type field
switch generic.Type {
case "TelemetryBasic":
    var telemetryBasic TelemetryBasic
    err = json.Unmarshal(generic.Data, &telemetryBasic)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid TelemetryBasic JSON",
            "cause": err.Error(),
        })
        return
    }

    // Write the basic telemetry to InfluxDB
    WriteTelemetry(telemetryBasic)

case "TelemetryFull":
    var telemetryFull TelemetryFull
    err = json.Unmarshal(generic.Data, &telemetryFull)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid TelemetryFull JSON",
            "cause": err.Error(),
        })
        return
    }

    // Write the full telemetry to InfluxDB
    WriteTelemetry(telemetryFull)

default:
    c.JSON(http.StatusBadRequest, gin.H{
        "error": "Unknown telemetry type",
    })
    return
}
}