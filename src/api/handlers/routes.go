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
	LocalDateTime             time.Time     `json:"localdatetime"`
	ExecutionDuration         time.Duration `json:"executionduration"`
	Failed                    bool          `json:"failed"`
	Exception                 string        `json:"exception"`
	Manufacturer              string        `json:"manufacturer"`
	Model                     string        `json:"model"`
	TotalPhysicalMemory       int           `json:"totalphysicalmemory"`
	NumberOfProcessors        int           `json:"numberofprocessors"`
	NumberOfLogicalProcessors int           `json:"numberoflogicalprocessors"`
	PartOfDomain              bool          `json:"partofdomain"`
	HardwareSerialNumber      string        `json:"hardwareserialnumber"`
	BootDriveSerial           string        `json:"bootdriveserial"`
	OSType                    string        `json:"ostype"`
	OSArchitecture            string        `json:"osarchitecture"`
	OSVersion                 string        `json:"version"`
	OSBuildNumber             string        `json:"buildnumber"`
	PowerShellVersion         string        `json:"powershellversion"`
	HostVersion               string        `json:"hostversion"`
	HostName                  string        `json:"hostname"`
	HostUI                    string        `json:"hostui"`
	HostCulture               string        `json:"hostculture"`
	HostUICulture             string        `json:"hostuiculture"`
	ModuleName                string        `json:"modulename"`
	ModuleVersion             string        `json:"moduleversion"`
	ModulePath                string        `json:"modulepath"`
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
