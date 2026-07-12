package port

import (
	"context"

	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/valueobject"
)

var (
	ErrEmployeeNotFound = xerr.New(xerr.CodeRecordNotFound,
		xerr.WithMeta("employee", xerr.ErrorReasonNotFound))
	ErrEmployeeAlreadyExists = xerr.New(xerr.CodeAlreadyExists,
		xerr.WithMeta("employee", xerr.ErrorReasonAlreadyExists))
	ErrEmployeeRepository = xerr.New(xerr.CodeDatabaseError)
)

type EmployeeRepository interface {
	Create(ctx context.Context, employee *entity.Employee) error
	Update(ctx context.Context, employee *entity.Employee) error
	Delete(ctx context.Context, id valueobject.EmployeeID, organizationID valueobject.OrganizationID) error

	FindByID(ctx context.Context, id valueobject.EmployeeID, organizationID valueobject.OrganizationID) (*entity.Employee, error)
	FindByUserIDAndOrganizationID(ctx context.Context, userID valueobject.UserID, organizationID valueobject.OrganizationID) (*entity.Employee, error)
	ListActiveByUserID(ctx context.Context, userID valueobject.UserID) ([]*entity.Employee, error)
	ListByOrganizationID(ctx context.Context, organizationID valueobject.OrganizationID) ([]*entity.Employee, error)
	CountActiveByRoleID(ctx context.Context, roleID valueobject.RoleID) (int64, error)
}
