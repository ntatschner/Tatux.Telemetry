package handlers

import (
	"context"
	"log"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/ntatschner/Tatux.Telemetry/src/api/system"
)

var (
	InfluxDBUrl     = system.GetEnv("INFLUXDB_URL", "", true) + ":" + system.GetEnv("INFLUXDB_PORT", "8086", false)
	InfluxDBToken   = system.GetEnv("INFLUXDB_TOKEN", "", true)
	INFLUXDB_BUCKET = system.GetEnv("INFLUXDB_BUCKET", "", true)
	INFLUXDB_ORG    = system.GetEnv("INFLUXDB_ORG", "", true)
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
	point := influxdb2.NewPointWithMeasurement(telemetry.ModuleName).
		AddField("id", telemetry.ID).
		AddTag("commandName", telemetry.CommandName).
		AddField("complete", telemetry.Complete).
		AddField("localDateTime", telemetry.LocalDateTime).
		AddField("executionDuration", int64(telemetry.ExecutionDuration)).
		AddField("failed", telemetry.Failed).
		AddField("exception", telemetry.Exception).
		AddField("manufacturer", telemetry.Manufacturer).
		AddField("model", telemetry.Model).
		AddField("totalPhysicalMemory", telemetry.TotalPhysicalMemory).
		AddField("numberOfProcessors", telemetry.NumberOfProcessors).
		AddField("numberOfLogicalProcessors", telemetry.NumberOfLogicalProcessors).
		AddField("partOfDomain", telemetry.PartOfDomain).
		AddField("hardwareSerialNumber", telemetry.HardwareSerialNumber).
		AddField("bootDriveSerial", telemetry.BootDriveSerial).
		AddField("osType", telemetry.OSType).
		AddField("osArchitecture", telemetry.OSArchitecture).
		AddField("osVersion", telemetry.OSVersion).
		AddField("osBuildNumber", telemetry.OSBuildNumber).
		AddField("powerShellVersion", telemetry.PowerShellVersion).
		AddField("hostVersion", telemetry.HostVersion).
		AddField("hostName", telemetry.HostName).
		AddField("hostUI", telemetry.HostUI).
		AddField("hostCulture", telemetry.HostCulture).
		AddField("hostUICulture", telemetry.HostUICulture).
		AddTag("moduleName", telemetry.ModuleName).
		AddField("moduleVersion", telemetry.ModuleVersion).
		AddField("modulePath", telemetry.ModulePath).
		SetTime(time.Now())

	// Write the point
	writeAPI := client.WriteAPIBlocking(INFLUXDB_ORG, INFLUXDB_BUCKET)
	writeAPI.WritePoint(context.Background(), point)
}
