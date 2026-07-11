package authorization

import (
	"github.com/TrueFlowDev/Backend/internal/module/authorization/application/usecase"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/presentation/http/controller"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"authorization",
	fx.Provide(
		usecase.NewListPermissionsUseCase,
		controller.NewListPermissionsController,
	),
	fx.Invoke(
		controller.RegisterListPermissionsController,
	),
)
