package user

import (
	"github.com/TrueFlowDev/Backend/internal/module/user/application/usecase"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/user/infrastructure/adapter"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"user",
	fx.Provide(
		fx.Annotate(
			adapter.NewUUIDGenerator,
			fx.As(new(port.UserIdGenerator)),
		),
		fx.Annotate(
			adapter.NewUserRepository,
			fx.As(new(port.UserRepository)),
		),
		usecase.NewRegisterUserUsecase,
		usecase.NewFindUserByPhoneUsecase,
	),
)
