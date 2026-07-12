package adapter

import (
	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/valueobject"
	"github.com/google/uuid"
)

var _ port.EmployeeIDGenerator = (*UUIDGenerator)(nil)

type UUIDGenerator struct{}

func NewUUIDGenerator() *UUIDGenerator { return &UUIDGenerator{} }

func (g *UUIDGenerator) Generate() valueobject.EmployeeID {
	return valueobject.NewEmployeeID(uuid.NewString())
}
