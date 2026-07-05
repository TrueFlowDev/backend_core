package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/user/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"
)

type FindUserByIDInput struct {
	ID string
}

type FindUserByIDUsecase struct {
	userRepository port.UserRepository
}

func NewFindUserByIDUsecase(
	userRepository port.UserRepository,
) *FindUserByIDUsecase {
	return &FindUserByIDUsecase{
		userRepository: userRepository,
	}
}

func (u *FindUserByIDUsecase) Execute(ctx context.Context, input FindUserByIDInput) (*entity.User, error) {
	userID := value_object.NewUserID(input.ID)

	user, err := u.userRepository.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
