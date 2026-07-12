package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/valueobject"
	sharedPort "github.com/TrueFlowDev/Backend/internal/shared/domain/port"
)

type CreateOrganizationWithOwnerInput struct {
	Category string
	Name     string

	OwnerUserID         string
	OwnerJobTitle       string
	OwnerEmploymentType string
}

type CreateOrganizationWithOwnerOutput struct {
	OrganizationID string
	RoleID         string
	EmployeeID     string
}

type CreateOrganizationWithOwnerUsecase struct {
	organizationIDGenerator port.OrganizationIDGenerator
	organizationRepository  port.OrganizationRepository
	roleCreator             port.RoleCreator
	employeeCreator         port.EmployeeCreator
	txManager               sharedPort.TxManager
}

func NewCreateOrganizationWithOwnerUsecase(
	organizationIDGenerator port.OrganizationIDGenerator,
	organizationRepository port.OrganizationRepository,
	roleCreator port.RoleCreator,
	employeeCreator port.EmployeeCreator,
	txManager sharedPort.TxManager,
) *CreateOrganizationWithOwnerUsecase {
	return &CreateOrganizationWithOwnerUsecase{
		organizationIDGenerator: organizationIDGenerator,
		organizationRepository:  organizationRepository,
		roleCreator:             roleCreator,
		employeeCreator:         employeeCreator,
		txManager:               txManager,
	}
}

func (u *CreateOrganizationWithOwnerUsecase) Execute(
	ctx context.Context,
	input CreateOrganizationWithOwnerInput,
) (CreateOrganizationWithOwnerOutput, error) {
	organizationCategory, err := valueobject.ParseOrganizationCategory(input.Category)
	if err != nil {
		return CreateOrganizationWithOwnerOutput{}, err
	}

	newOrganizationID := u.organizationIDGenerator.Generate()

	newOrganization, err := entity.NewOrganization(newOrganizationID, input.Name, organizationCategory)
	if err != nil {
		return CreateOrganizationWithOwnerOutput{}, err
	}

	var roleID, employeeID string

	err = u.txManager.WithinTx(ctx, func(ctx context.Context) error {
		if err := u.organizationRepository.Create(ctx, newOrganization); err != nil {
			return err
		}

		roleOutput, err := u.roleCreator.CreateOwnerRole(ctx, port.RoleCreatorInput{
			OrganizationID: newOrganizationID.Value(),
		})
		if err != nil {
			return err
		}
		roleID = roleOutput.ID

		employeeOutput, err := u.employeeCreator.Create(ctx, port.EmployeeCreatorInput{
			UserID:           input.OwnerUserID,
			OrganizationID:   newOrganizationID.Value(),
			RoleID:           roleID,
			JobTitle:         input.OwnerJobTitle,
			MembershipStatus: "active",
			EmploymentType:   input.OwnerEmploymentType,
		})
		if err != nil {
			return err
		}
		employeeID = employeeOutput.ID

		return nil
	})
	if err != nil {
		return CreateOrganizationWithOwnerOutput{}, err
	}

	return CreateOrganizationWithOwnerOutput{
		OrganizationID: newOrganizationID.Value(),
		RoleID:         roleID,
		EmployeeID:     employeeID,
	}, nil
}
