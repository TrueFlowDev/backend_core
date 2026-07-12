package organization

import (
	"github.com/TrueFlowDev/Backend/internal/module/organization/application/usecase"
	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/organization/infrastructure/adapter"
	"github.com/TrueFlowDev/Backend/internal/module/organization/presentation/http/controller"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"organization",
	fx.Provide(
		fx.Annotate(
			adapter.NewOrganizationRepository,
			fx.As(new(port.OrganizationRepository)),
		),
		fx.Annotate(
			adapter.NewUUIDGenerator,
			fx.As(new(port.OrganizationIDGenerator)),
		),
		fx.Annotate(
			adapter.NewRoleCreator,
			fx.As(new(port.RoleCreator)),
		),
		fx.Annotate(
			adapter.NewEmployeeCreator,
			fx.As(new(port.EmployeeCreator)),
		),
		usecase.NewCreateOrganizationWithOwnerUsecase,
		usecase.NewFindOrganizationByIDUsecase,
		usecase.NewFindOrganizationsByIDsUsecase,
		usecase.NewUpdateOrganizationUsecase,
		controller.NewCreateOrganizationController,
		controller.NewUpdateOrganizationController,
	),
	fx.Invoke(
		controller.RegisterCreateOrganizationController,
		controller.RegisterUpdateOrganizationController,
	),
)
