package app

import (
	"log"

	ginrouter "github.com/gin-gonic/gin"
	"github.com/ntatschner/Tatux.Telemetry/src/api/handlers"
	"github.com/ntatschner/Tatux.Telemetry/src/api/system"
)

var (
	InfluxDBUrl   = system.GetEnv("INFLUXDB_URL", "", false) + ":" + system.GetEnv("INFLUXDB_PORT", "", false)
	InfluxDBToken = system.GetEnv("INFLUXDB_TOKEN", "", false)
	listenPort    = system.GetEnv("LISTENONPORT", "9000", false)
)

func Start() {
	// Starting the router
	router := ginrouter.Default()

	// Routes
	router.GET("/ping", handlers.GetPing)
	router.PUT("/api/telemetry", handlers.PutTelemetry)
	router.GET("/api/health", handlers.GetSystemHealth)

	// Start server
	log.Println("Starting API server on :" + listenPort)
	router.Run(":" + listenPort)

	// Connect to InfluxDB
	handlers.ConnectInfluxDB(InfluxDBUrl, InfluxDBToken)
}
