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
		fx.Annotate(
			adapter.NewRoleRepository,
			fx.As(new(port.RoleRepository)),
		),
		usecase.NewListPermissionsUseCase,
		usecase.NewCreateRoleUsecase,
		controller.NewListPermissionsController,
	),
	fx.Invoke(
		controller.RegisterListPermissionsController,
	),
)
