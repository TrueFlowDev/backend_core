package adapter

import (
	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/valueobject"
	"github.com/google/uuid"
)

var _ port.OrganizationIDGenerator = (*UUIDGenerator)(nil)

type UUIDGenerator struct{}

func NewUUIDGenerator() *UUIDGenerator { return &UUIDGenerator{} }

func (g *UUIDGenerator) Generate() valueobject.OrganizationID {
	return valueobject.NewOrganizationID(uuid.NewString())
}
