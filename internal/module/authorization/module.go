package authorization

import (
	"github.com/TrueFlowDev/Backend/internal/module/authorization/application/usecase"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/infrastructure/adapter"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/presentation/http/controller"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"authorization",
	fx.Provide(
		fx.As(
			adapter.NewRoleRepository,
			fx.As(new(port.RoleRepository)),
		),
		usecase.NewListPermissionsUseCase,
		controller.NewListPermissionsController,
	),
	fx.Invoke(
		controller.RegisterListPermissionsController,
	),
)
