package port

import "github.com/TrueFlowDev/Backend/internal/module/employee/domain/valueobject"

type EmployeeIDGenerator interface {
	Generate() valueobject.EmployeeID
}
