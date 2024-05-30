package identity

import (
	"fmt"
	"github.com/ntatschner/Tatux.Telemetry/src/telemetry/pkg/telemetry/domain"
	"github.com/google/uuid"
)

type UUIDGenerator struct {}

func (g *UUIDGenerator) NewID() string {
    return uuid.New().String()
}
var _ domain.IDGenerator = &UUIDGenerator{}