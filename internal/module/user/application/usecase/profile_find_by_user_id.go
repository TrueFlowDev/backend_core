package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/user/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"
)

type FindProfileByUserIDInput struct {
	UserID string
}

type FindProfileByUserIDUsecase struct {
	profileRepository port.ProfileRepository
}

func NewFindProfileByUserIDUsecase(
	profileRepository port.ProfileRepository,
) *FindProfileByUserIDUsecase {
	return &FindProfileByUserIDUsecase{
		profileRepository: profileRepository,
	}
}

func (u *FindProfileByUserIDUsecase) Execute(ctx context.Context, input FindProfileByUserIDInput) (*entity.Profile, error) {
	userID := value_object.NewUserID(input.UserID)

	profile, err := u.profileRepository.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return profile, nil
}
