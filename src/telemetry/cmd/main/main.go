package main

import (
	"github.com/ntatschner/Tatux.Telemetry/src/telemetry/pkg/logging"
	"github.com/ntatschner/Tatux.Telemetry/src/telemetry/pkg/telemetry/adapters"
	"github.com/ntatschner/Tatux.Telemetry/src/telemetry/pkg/telemetry/domain"
)

func main() {
	logger := logging.NewStdLogger()

	telemetryService := domain.NewTelemetryService(logger)

	adapter := adapters.NewTelemetryAdapter(telemetryService, logger)

	logger.Info("Starting telemetry service")

}
