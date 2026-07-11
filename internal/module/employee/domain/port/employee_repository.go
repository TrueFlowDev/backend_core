package port

import (
	"context"

	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/value_object"
)

var (
	ErrEmployeeNotFound   = xerr.New(xerr.CodeRecordNotFound, xerr.WithMeta("employee", xerr.ErrorReasonNotFound))
	ErrEmployeeRepository = xerr.New(xerr.CodeDatabaseError)
)

type EmployeeRepository interface {
	Create(ctx context.Context, employee *entity.Employee) error

	FindByID(ctx context.Context, id value_object.EmployeeID) (*entity.Employee, error)
}
