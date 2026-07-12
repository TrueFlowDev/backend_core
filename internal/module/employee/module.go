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
		usecase.NewCreateEmployeeUsecase,
		usecase.NewListMyDashboardsUsecase,
		usecase.NewFindEmployeeRoleUsecase,
		usecase.NewCountActiveEmployeesByRoleUsecase,
		controller.NewListMyDashboardsController,
	),
	fx.Invoke(
		controller.RegisterListMyDashboardsController,
	),
)
