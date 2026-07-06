package adapter

import (
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"
	"github.com/google/uuid"
)

var _ port.UserIdGenerator = (*UUIDGenerator)(nil)

type UUIDGenerator struct{}

func NewUUIDGenerator() *UUIDGenerator { return &UUIDGenerator{} }

func (g *UUIDGenerator) Generate() value_object.UserID {
	return value_object.NewUserID(uuid.NewString())
}
