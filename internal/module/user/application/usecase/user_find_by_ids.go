package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/user/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/valueobject"
)

type FindUsersByIDsInput struct {
	IDs []string
}

type FindUsersByIDsUsecase struct {
	userRepository port.UserRepository
}

func NewFindUsersByIDsUsecase(userRepository port.UserRepository) *FindUsersByIDsUsecase {
	return &FindUsersByIDsUsecase{userRepository: userRepository}
}

func (u *FindUsersByIDsUsecase) Execute(ctx context.Context, input FindUsersByIDsInput) ([]*entity.User, error) {
	ids := make([]valueobject.UserID, len(input.IDs))
	for i, raw := range input.IDs {
		ids[i] = valueobject.NewUserID(raw)
	}

	return u.userRepository.FindByIDs(ctx, ids)
}
