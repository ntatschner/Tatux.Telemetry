package incoming

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ntatschner/Tatux.Telemetry/src/telemetry/pkg/logging"
	"github.com/ntatschner/Tatux.Telemetry/src/telemetry/pkg/telemetry/domain"
)

var logger = logging.NewStdLogger()
// HandleJSONPayload takes a raw JSON payload and determines whether it should be
// unmarshalled into a PowerShellData or PipelineData struct.
func HandleJSONPayload(c *gin.Context) error {
	var BaseData domain.BaseAPIPayload
	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error reading request body",
		})
		logger.Error("Error reading request body: %t", err)
		return err
	}

	err = json.Unmarshal(payload, &BaseData)
	if err != nil {
		logger.Error("Error unmarshalling payload: %t", err)
		return err
	}
	PayloadTest, err := json.Marshal(BaseData.Payload)
	if err != nil {
		logger.Error("Error marshalling payload: %t", err)
		return err
	}

	switch BaseData.SourceType {
	case "PowerShellData":
		var data domain.PowerShellData
		err := json.Unmarshal(PayloadTest, &data)
		if err != nil {
			logger.Error("Error unmarshalling PowerShellData: %t", err)
			return err
		}
		
	case "PipelineData":
		var data domain.PipelineData
		err := json.Unmarshal(payload, &data)
		if err != nil {
			logger.Error("Error unmarshalling PipelineData: %t", err)
			return err
		}
		return &data, nil
	default:
		logger.Error("Unknown data type: " + dataType)
		return nil, errors.New("unknown data type")
	}
}
