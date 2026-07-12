package employee

import (
	"github.com/TrueFlowDev/Backend/internal/module/employee/application/usecase"
	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/employee/infrastructure/adapter"
	"github.com/TrueFlowDev/Backend/internal/module/employee/presentation/http/controller"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"employee",
	fx.Provide(
		fx.Annotate(
			adapter.NewEmployeeRepository,
			fx.As(new(port.EmployeeRepository)),
		),
		fx.Annotate(
			adapter.NewUUIDGenerator,
			fx.As(new(port.EmployeeIDGenerator)),
		),
		fx.Annotate(
			adapter.NewOrganizationFinder,
			fx.As(new(port.OrganizationFinder)),
		),
		fx.Annotate(
			adapter.NewRoleFinder,
			fx.As(new(port.RoleFinder)),
		),
		fx.Annotate(
			adapter.NewUserFinder,
			fx.As(new(port.UserFinder)),
		),
		usecase.NewCreateEmployeeUsecase,
		usecase.NewListMyDashboardsUsecase,
		usecase.NewFindEmployeeRoleUsecase,
		usecase.NewCountActiveEmployeesByRoleUsecase,
		usecase.NewAddEmployeeUsecase,
		usecase.NewListEmployeesUsecase,
		usecase.NewUpdateEmployeeUsecase,
		usecase.NewRemoveEmployeeUsecase,
		controller.NewListMyDashboardsController,
		controller.NewAddEmployeeController,
		controller.NewListEmployeesController,
		controller.NewRemoveEmployeeController,
		controller.NewUpdateEmployeeController,
	),
	fx.Invoke(
		controller.RegisterListMyDashboardsController,
		controller.RegisterAddEmployeeController,
		controller.RegisterListEmployeesController,
		controller.RegisterRemoveEmployeeController,
		controller.RegisterUpdateEmployeeController,
	),
)
