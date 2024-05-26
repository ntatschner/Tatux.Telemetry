package incoming

import (
	"encoding/json"
	"errors"
)

// HandleJSONPayload takes a raw JSON payload and determines whether it should be
// unmarshalled into a PowerShellData or PipelineData struct.
func HandleJSONPayload(payload []byte, dataType string) (interface{}, error) {
	switch dataType {
	case "PowerShellData":
		var data PowerShellData
		err := json.Unmarshal(payload, &data)
		if err != nil {
			return nil, err
		}
		return &data, nil
	case "PipelineData":
		var data PipelineData
		err := json.Unmarshal(payload, &data)
		if err != nil {
			return nil, err
		}
		return &data, nil
	default:
		return nil, errors.New("unknown data type")
	}
}
