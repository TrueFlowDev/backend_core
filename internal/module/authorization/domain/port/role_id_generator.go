package port

import "github.com/TrueFlowDev/Backend/internal/module/authorization/domain/valueobject"

type EmployeeIDGenerator interface {
	Generate() valueobject.RoleID
}
