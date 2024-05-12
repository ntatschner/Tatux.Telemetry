package app

import (
	"log"
	"strings"

	ginrouter "github.com/gin-gonic/gin"
	"github.com/ntatschner/Tatux.Telemetry/src/api/handlers"
	"github.com/ntatschner/Tatux.Telemetry/src/api/system"
)

var (
	InfluxDBUrl    = system.GetEnv("INFLUXDB_URL", "", true) + ":" + system.GetEnv("INFLUXDB_PORT", "8086", false)
	InfluxDBToken  = system.GetEnv("INFLUXDB_TOKEN", "", true)
	listenPort     = system.GetEnv("LISTENONPORT", "9000", false)
	TrustedProxies = system.GetEnv("TRUSTEDPROXIES", "127.0.0.1", false)
)

func Start() {
	// Starting the router
	router := ginrouter.Default()
	router.SetTrustedProxies(strings.Split(TrustedProxies, ","))
	// Routes
	router.GET("/ping", handlers.GetPing)
	router.PUT("/api/telemetry", handlers.PutTelemetry)
	router.GET("/api/health", handlers.GetSystemHealth)

	// Start server
	log.Println("Starting API server on :" + listenPort)
	router.Run(":" + listenPort)

	// Connect to InfluxDB
	handlers.ConnectInfluxDB(InfluxDBUrl, InfluxDBToken)
	system.GetGeoLocationDatabase()
}
