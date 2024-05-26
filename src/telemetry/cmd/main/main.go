package main

import (
	"github.com/ntatschner/Tatux.Telemetry/src/telemetry/pkg/logging"
)

func main() {
	logger := logging.NewStdLogger()

	logger.Info("Starting telemetry service")

}
