package main

import (
	"telemetry/pkg/logging"
	"telemetry/pkg/telemetry/adapters"
	"telemetry/pkg/telemetry/domain"
)

func main() {
	logger := logging.NewStdLogger()

	telemetryService := domain.NewTelemetryService(logger)

	adapter := adapters.NewTelemetryAdapter(telemetryService, logger)

	logger.Info("Starting telemetry service")

}
