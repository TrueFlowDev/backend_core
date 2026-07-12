package port

import (
	"context"

	"github.com/Ali127Dev/xerr"
)

var ErrRoleInUse = xerr.New(xerr.CodeBadRequest, xerr.WithMeta("role", xerr.ErrorReasonInvalidFormat))

type EmployeeRoleUsageChecker interface {
	CountActiveEmployeesByRole(ctx context.Context, roleID string) (int64, error)
}
