package adapter

import (
	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/valueobject"
	"github.com/google/uuid"
)

var _ port.RoleIDGenerator = (*UUIDGenerator)(nil)

type UUIDGenerator struct{}

func NewUUIDGenerator() *UUIDGenerator { return &UUIDGenerator{} }

func (g *UUIDGenerator) Generate() valueobject.RoleID {
	return valueobject.NewRoleID(uuid.NewString())
}
