package adapter

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/port"
	userUsecase "github.com/TrueFlowDev/Backend/internal/module/user/application/usecase"
)

var _ port.UserFinder = (*UserFinder)(nil)

type UserFinder struct {
	findUserByPhoneUsecase       *userUsecase.FindUserByPhoneUsecase
	findUsersByIDsUsecase        *userUsecase.FindUsersByIDsUsecase
	findProfilesByUserIDsUsecase *userUsecase.FindProfilesByUserIDsUsecase
}

func NewUserFinder(
	findUserByPhoneUsecase *userUsecase.FindUserByPhoneUsecase,
	findUsersByIDsUsecase *userUsecase.FindUsersByIDsUsecase,
	findProfilesByUserIDsUsecase *userUsecase.FindProfilesByUserIDsUsecase,
) *UserFinder {
	return &UserFinder{
		findUserByPhoneUsecase:       findUserByPhoneUsecase,
		findUsersByIDsUsecase:        findUsersByIDsUsecase,
		findProfilesByUserIDsUsecase: findProfilesByUserIDsUsecase,
	}
}

func (a *UserFinder) FindByPhone(ctx context.Context, phone string) (port.UserFinderOutput, error) {
	user, err := a.findUserByPhoneUsecase.Execute(ctx, userUsecase.FindUserByPhoneInput{Phone: phone})
	if err != nil {
		return port.UserFinderOutput{}, err
	}

	return port.UserFinderOutput{ID: user.ID().Value(), Phone: user.Phone().Value()}, nil
}

func (a *UserFinder) FindByIDs(ctx context.Context, userIDs []string) ([]port.UserFinderOutput, error) {
	if len(userIDs) == 0 {
		return nil, nil
	}

	users, err := a.findUsersByIDsUsecase.Execute(ctx, userUsecase.FindUsersByIDsInput{IDs: userIDs})
	if err != nil {
		return nil, err
	}

	profiles, err := a.findProfilesByUserIDsUsecase.Execute(ctx, userUsecase.FindProfilesByUserIDsInput{UserIDs: userIDs})
	if err != nil {
		return nil, err
	}

	profileByUserID := make(map[string]struct {
		FirstName string
		LastName  string
	}, len(profiles))
	for _, p := range profiles {
		profileByUserID[p.UserID().Value()] = struct {
			FirstName string
			LastName  string
		}{FirstName: p.FirstName(), LastName: p.LastName()}
	}

	result := make([]port.UserFinderOutput, len(users))
	for i, u := range users {
		output := port.UserFinderOutput{ID: u.ID().Value(), Phone: u.Phone().Value()}
		if profile, ok := profileByUserID[u.ID().Value()]; ok {
			output.FirstName = profile.FirstName
			output.LastName = profile.LastName
		}
		result[i] = output
	}

	return result, nil
}
