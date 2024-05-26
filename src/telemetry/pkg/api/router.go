package api

import (
	ginrouter "github.com/gin-gonic/gin"

	"github.com/ntatschner/Tatux.Telemetry/src/telemetry/logging"
	"github.com/ntatschner/Tatux.Telemetry/src/telemetry/pkg/telemetry/ports/incoming"
)

func NewRouter() *ginrouter.Engine {
    router := ginrouter.Default()

    // Define routes
    router.GET("/api/ping", HealthCheck)
	router.POST("/api/telemetry", incoming.HandleJSONPayload)

    return router
}