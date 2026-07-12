package usecase

import (
	"context"
	"errors"

	"github.com/TrueFlowDev/Backend/internal/module/user/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/valueobject"
)

type SaveProfileInput struct {
	UserID    string
	Email     string
	FirstName string
	LastName  string
	Headline  string
	Bio       string
}

type SaveProfileOutput struct {
	UserID string
}

type SaveProfileUsecase struct {
	profileRepository port.ProfileRepository
}

func NewSaveProfileUsecase(
	profileRepository port.ProfileRepository,
) *SaveProfileUsecase {
	return &SaveProfileUsecase{
		profileRepository: profileRepository,
	}
}

func (u *SaveProfileUsecase) Execute(
	ctx context.Context,
	input SaveProfileInput,
) (SaveProfileOutput, error) {
	userID := valueobject.NewUserID(input.UserID)

	email, err := valueobject.NewEmail(input.Email)
	if err != nil {
		return SaveProfileOutput{}, err
	}

	profile, err := u.profileRepository.FindByUserID(ctx, userID)
	if err != nil {
		if !errors.Is(err, port.ErrProfileNotFound) {
			return SaveProfileOutput{}, err
		}

		profile, err = entity.NewProfile(
			userID,
			email,
			input.FirstName,
			input.LastName,
		)
		if err != nil {
			return SaveProfileOutput{}, err
		}
	} else {
		if err := profile.UpdateFirstName(input.FirstName); err != nil {
			return SaveProfileOutput{}, err
		}

		if err := profile.UpdateLastName(input.LastName); err != nil {
			return SaveProfileOutput{}, err
		}

		profile.UpdateEmail(email)
	}

	if err := profile.UpdateHeadline(input.Headline); err != nil {
		return SaveProfileOutput{}, err
	}

	if err := profile.UpdateBio(input.Bio); err != nil {
		return SaveProfileOutput{}, err
	}

	if err := u.profileRepository.Save(ctx, profile); err != nil {
		return SaveProfileOutput{}, err
	}

	return SaveProfileOutput{
		UserID: profile.UserID().Value(),
	}, nil
}
