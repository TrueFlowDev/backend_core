package user

import (
	"github.com/TrueFlowDev/Backend/internal/module/user/applicaiton/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"user",
	fx.Provide(
		usecase.NewRegisterUserUsecase,
	),
)
