package adapter

import (
	"context"
	"errors"

	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/valueobject"
	"github.com/TrueFlowDev/Backend/internal/module/employee/infrastructure/dao"
	"github.com/TrueFlowDev/Backend/internal/module/employee/infrastructure/mapper"
	"github.com/TrueFlowDev/Backend/internal/shared/infrastructure/database"
	"gorm.io/gorm"
)

var _ port.EmployeeRepository = (*EmployeeRepository)(nil)

type EmployeeRepository struct {
	*database.BaseRepo
}

func NewEmployeeRepository(base *database.BaseRepo) *EmployeeRepository {
	return &EmployeeRepository{BaseRepo: base}
}

func (r *EmployeeRepository) Create(ctx context.Context, employee *entity.Employee) error {
	q := dao.Use(r.Executor(ctx))

	mappedEmployee := mapper.EmployeeEntityToModel(employee)
	if err := q.WithContext(ctx).Employee.Create(mappedEmployee); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return port.ErrEmployeeAlreadyExists
		}
		return xerr.Wrap(err, port.ErrEmployeeRepository.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "employee_create"))
	}

	return nil
}

func (r *EmployeeRepository) FindByID(
	ctx context.Context, id valueobject.EmployeeID,
	organizationID valueobject.OrganizationID,
) (*entity.Employee, error) {
	q := dao.Use(r.Executor(ctx))

	model, err := q.WithContext(ctx).Employee.
		Where(
			q.Employee.ID.Eq(id.Value()),
			q.Employee.OrganizationID.Eq(organizationID.Value()),
		).
		First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, port.ErrEmployeeNotFound
		}
		return nil, xerr.Wrap(err, port.ErrEmployeeRepository.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "employee_find_by_id"))
	}
	mappedUser, err := mapper.EmployeeModelToEntity(model)
	if err != nil {
		return nil, err
	}
	return mappedUser, nil
}

func (r *EmployeeRepository) ListActiveByUserID(
	ctx context.Context, userID valueobject.UserID,
) ([]*entity.Employee, error) {
	q := dao.Use(r.Executor(ctx))

	models, err := q.WithContext(ctx).Employee.
		Where(
			q.Employee.UserID.Eq(userID.Value()),
			q.Employee.MembershipStatus.Eq(valueobject.MembershipStatusActive.Value()),
		).
		Find()
	if err != nil {
		return nil, xerr.Wrap(err, port.ErrEmployeeRepository.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "employee_list_active_by_user_id"))
	}

	employees := make([]*entity.Employee, 0, len(models))
	for _, m := range models {
		e, err := mapper.EmployeeModelToEntity(m)
		if err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}

	return employees, nil
}
