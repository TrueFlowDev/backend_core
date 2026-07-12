package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/user/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/valueobject"
)

type FindProfilesByUserIDsInput struct {
	UserIDs []string
}

type FindProfilesByUserIDsUsecase struct {
	profileRepository port.ProfileRepository
}

func NewFindProfilesByUserIDsUsecase(profileRepository port.ProfileRepository) *FindProfilesByUserIDsUsecase {
	return &FindProfilesByUserIDsUsecase{profileRepository: profileRepository}
}

func (u *FindProfilesByUserIDsUsecase) Execute(ctx context.Context, input FindProfilesByUserIDsInput) ([]*entity.Profile, error) {
	ids := make([]valueobject.UserID, len(input.UserIDs))
	for i, raw := range input.UserIDs {
		ids[i] = valueobject.NewUserID(raw)
	}

	return u.profileRepository.FindByUserIDs(ctx, ids)
}
