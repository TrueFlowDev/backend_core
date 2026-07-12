package adapter

import (
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/valueobject"
	"github.com/google/uuid"
)

var _ port.UserIDGenerator = (*UUIDGenerator)(nil)

type UUIDGenerator struct{}

func NewUUIDGenerator() *UUIDGenerator { return &UUIDGenerator{} }

func (g *UUIDGenerator) Generate() valueobject.UserID {
	return valueobject.NewUserID(uuid.NewString())
}
