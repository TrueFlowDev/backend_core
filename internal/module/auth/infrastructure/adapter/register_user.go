package adapter

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/port"
	userUsecase "github.com/TrueFlowDev/Backend/internal/module/user/application/usecase"
)

var _ port.UserRegisterer = (*UserRegisterer)(nil)

type UserRegisterer struct {
	registerUserUsecase *userUsecase.RegisterUserUsecase
}

func NewUserRegisterer(
	registerUserUsecase *userUsecase.RegisterUserUsecase,
) *UserRegisterer {
	return &UserRegisterer{registerUserUsecase: registerUserUsecase}
}

func (a *UserRegisterer) Register(ctx context.Context, input port.UserRegistererInput) (port.UserRegistererOutput, error) {
	output, err := a.registerUserUsecase.Execute(ctx, userUsecase.RegisterUserInput{
		Phone:          input.Phone,
		HashedPassword: input.HashedPassword,
	})
	if err != nil {
		return port.UserRegistererOutput{}, err
	}
	return port.UserRegistererOutput{
		ID: output.ID,
	}, err
}
