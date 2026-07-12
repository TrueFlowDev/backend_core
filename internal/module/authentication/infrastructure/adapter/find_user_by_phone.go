package adapter

import (
	"context"
	"errors"

	"github.com/TrueFlowDev/Backend/internal/module/authentication/domain/port"
	userUsecase "github.com/TrueFlowDev/Backend/internal/module/user/application/usecase"
	user_port "github.com/TrueFlowDev/Backend/internal/module/user/domain/port"
)

var _ port.UserFinderByPhone = (*UserFinderByPhone)(nil)

type UserFinderByPhone struct {
	usecase *userUsecase.FindUserByPhoneUsecase
}

func NewUserFinderByPhone(
	usecase *userUsecase.FindUserByPhoneUsecase,
) *UserFinderByPhone {
	return &UserFinderByPhone{usecase: usecase}
}

func (a *UserFinderByPhone) FindByPhone(ctx context.Context, input port.UserFinderByPhoneInput) (port.UserFinderByPhoneOutput, error) {
	output, err := a.usecase.Execute(ctx, userUsecase.FindUserByPhoneInput{
		Phone: input.Phone,
	})
	if errors.Is(err, user_port.ErrUserNotFound) {
		return port.UserFinderByPhoneOutput{}, port.ErrUserNotFound
	}
	if err != nil {
		return port.UserFinderByPhoneOutput{}, err
	}
	return port.UserFinderByPhoneOutput{
		ID:             output.ID().Value(),
		HashedPassword: output.Password().Value(),
	}, nil
}
