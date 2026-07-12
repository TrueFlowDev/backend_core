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

func (r *EmployeeRepository) CountActiveByRoleID(
	ctx context.Context, roleID valueobject.RoleID,
) (int64, error) {
	q := dao.Use(r.Executor(ctx))

	count, err := q.WithContext(ctx).Employee.
		Where(
			q.Employee.RoleID.Eq(roleID.Value()),
			q.Employee.MembershipStatus.Eq(valueobject.MembershipStatusActive.Value()),
		).
		Count()
	if err != nil {
		return 0, xerr.Wrap(err, port.ErrEmployeeRepository.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "employee_count_active_by_role_id"))
	}

	return count, nil
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

func (r *EmployeeRepository) FindByUserIDAndOrganizationID(
	ctx context.Context, userID valueobject.UserID, organizationID valueobject.OrganizationID,
) (*entity.Employee, error) {
	q := dao.Use(r.Executor(ctx))

	model, err := q.WithContext(ctx).Employee.
		Where(
			q.Employee.UserID.Eq(userID.Value()),
			q.Employee.OrganizationID.Eq(organizationID.Value()),
		).
		First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, port.ErrEmployeeNotFound
		}
		return nil, xerr.Wrap(err, port.ErrEmployeeRepository.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "employee_find_by_user_and_organization"))
	}

	return mapper.EmployeeModelToEntity(model)
}

func (r *EmployeeRepository) Update(ctx context.Context, employee *entity.Employee) error {
	q := dao.Use(r.Executor(ctx))

	employeeModel := mapper.EmployeeEntityToModel(employee)

	result, err := q.WithContext(ctx).Employee.
		Where(
			q.Employee.ID.Eq(employeeModel.ID),
			q.Employee.OrganizationID.Eq(employeeModel.OrganizationID),
		).
		Updates(employeeModel)
	if err != nil {
		return xerr.Wrap(err, port.ErrEmployeeRepository.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "employee_update"))
	}
	if result.RowsAffected == 0 {
		return port.ErrEmployeeNotFound
	}

	return nil
}

func (r *EmployeeRepository) Delete(
	ctx context.Context, id valueobject.EmployeeID, organizationID valueobject.OrganizationID,
) error {
	q := dao.Use(r.Executor(ctx))

	result, err := q.WithContext(ctx).Employee.
		Where(
			q.Employee.ID.Eq(id.Value()),
			q.Employee.OrganizationID.Eq(organizationID.Value()),
		).
		Delete()
	if err != nil {
		return xerr.Wrap(err, port.ErrEmployeeRepository.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "employee_delete"))
	}
	if result.RowsAffected == 0 {
		return port.ErrEmployeeNotFound
	}

	return nil
}

func (r *EmployeeRepository) ListByOrganizationID(
	ctx context.Context, organizationID valueobject.OrganizationID,
) ([]*entity.Employee, error) {
	q := dao.Use(r.Executor(ctx))

	models, err := q.WithContext(ctx).Employee.
		Where(q.Employee.OrganizationID.Eq(organizationID.Value())).
		Find()
	if err != nil {
		return nil, xerr.Wrap(err, port.ErrEmployeeRepository.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "employee_list_by_organization_id"))
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
