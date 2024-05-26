package incoming

import (
	"encoding/json"
	"errors"
	
	"github.com/ntatschner/Tatux.Telemetry/src/telemetry/pkg/telemetry/domain"
	"github.com/ntatschner/Tatux.Telemetry/src/telemetry/pkg/logging"
)

var logger = logging.NewStdLogger()
// HandleJSONPayload takes a raw JSON payload and determines whether it should be
// unmarshalled into a PowerShellData or PipelineData struct.
func HandleJSONPayload(payload []byte, dataType string) (interface{}, error) {
	switch dataType {
	case "PowerShellData":
		var data domain.PowerShellData
		err := json.Unmarshal(payload, &data)
		if err != nil {
			logger.Error("Error unmarshalling PowerShellData: ", err)
			return nil, err
		}
		return &data, nil
	case "PipelineData":
		var data domain.PipelineData
		err := json.Unmarshal(payload, &data)
		if err != nil {
			logger.Error("Error unmarshalling PipelineData: ", err)
			return nil, err
		}
		return &data, nil
	default:
		logger.Error("Unknown data type: " + dataType)
		return nil, errors.New("unknown data type")
	}
}
