package api

import (
	"github.com/gin-gonic/gin"

	"github.com/ntatschner/Tatux.Telemetry/src/telemetry/pkg/logging"
)
var logger = logging.NewStdLogger()

func HealthCheck(c *gin.Context) {
	logger.Info("Health check request received")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// 