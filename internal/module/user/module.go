package user

import (
	"github.com/TrueFlowDev/Backend/internal/module/user/application/usecase"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/user/infrastructure/adapter"
	"github.com/TrueFlowDev/Backend/internal/module/user/presentation/http/controller"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"user",
	fx.Provide(
		fx.Annotate(
			adapter.NewUUIDGenerator,
			fx.As(new(port.UserIDGenerator)),
		),
		fx.Annotate(
			adapter.NewUserRepository,
			fx.As(new(port.UserRepository)),
		),
		fx.Annotate(
			adapter.NewProfileRepository,
			fx.As(new(port.ProfileRepository)),
		),
		usecase.NewRegisterUserUsecase,
		usecase.NewFindUserByPhoneUsecase,
		usecase.NewFindUserByIDUsecase,
		usecase.NewFindProfileByUserIDUsecase,
		usecase.NewSaveProfileUsecase,
		usecase.NewFindUsersByIDsUsecase,
		usecase.NewFindProfilesByUserIDsUsecase,
		controller.NewGetMeController,
		controller.NewGetMyProfileController,
		controller.NewSaveMyProfileController,
	),
	fx.Invoke(
		controller.RegisterGetMeController,
		controller.RegisterGetMyProfileController,
		controller.RegisterSaveMyProfileController,
	),
)
