package authorization

import (
	"github.com/TrueFlowDev/Backend/internal/module/authorization/application/usecase"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/infrastructure/adapter"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/presentation/http/controller"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/presentation/http/middleware"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"authorization",
	fx.Provide(
		fx.Annotate(
			adapter.NewRoleRepository,
			fx.As(new(port.RoleRepository)),
		),
		fx.Annotate(
			adapter.NewUUIDGenerator,
			fx.As(new(port.RoleIDGenerator)),
		),
		fx.Annotate(
			adapter.NewEmployeeRoleFinder,
			fx.As(new(port.EmployeeRoleFinder)),
		),
		usecase.NewListPermissionsUseCase,
		usecase.NewCreateRoleUsecase,
		usecase.NewFindRoleByIDUsecase,
		usecase.NewCreateOwnerRoleUsecase,
		usecase.NewHasPermissionUsecase,
		usecase.NewListRolesUsecase,
		usecase.NewUpdateRoleUsecase,
		usecase.NewDeleteRoleUsecase,
		controller.NewListPermissionsController,
		controller.NewCreateRoleController,
		controller.NewListRolesController,
		controller.NewUpdateRoleController,
		controller.NewDeleteRoleController,
		middleware.NewPermissionGuard,
	),
	fx.Invoke(
		controller.RegisterListPermissionsController,
		controller.RegisterCreateRoleController,
		controller.RegisterListRolesController,
		controller.RegisterUpdateRoleController,
		controller.RegisterDeleteRoleController,
	),
)
