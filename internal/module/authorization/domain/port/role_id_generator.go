package port

import "github.com/TrueFlowDev/Backend/internal/module/authorization/domain/valueobject"

type RoleIDGenerator interface {
	Generate() valueobject.RoleID
}
