package port

import "github.com/TrueFlowDev/Backend/internal/module/employee/domain/value_object"

type EmployeeIDGenerator interface {
	Generate() value_object.EmployeeID
}
