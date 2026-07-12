package port

import (
	"context"

	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/valueobject"
)

var (
	ErrRoleNotFound      = xerr.New(xerr.CodeRecordNotFound, xerr.WithMeta("role", xerr.ErrorReasonNotFound))
	ErrRoleAlreadyExists = xerr.New(xerr.CodeAlreadyExists, xerr.WithMeta("role", xerr.ErrorReasonAlreadyExists))
	ErrRoleRepository    = xerr.New(xerr.CodeDatabaseError)
)

type RoleRepository interface {
	Create(ctx context.Context, role *entity.Role) error

	FindByID(ctx context.Context, id valueobject.RoleID, organizationID valueobject.OrganizationID) (*entity.Role, error)
}
