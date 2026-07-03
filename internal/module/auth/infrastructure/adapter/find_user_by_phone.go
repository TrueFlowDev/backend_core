package adapter

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/port"
	userUsecase "github.com/TrueFlowDev/Backend/internal/module/user/application/usecase"
)

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
	if err != nil {
		return port.UserFinderByPhoneOutput{}, err
	}
	return port.UserFinderByPhoneOutput{
		ID:             output.ID().Value(),
		HashedPassword: output.Password().Value(),
	}, err
}
