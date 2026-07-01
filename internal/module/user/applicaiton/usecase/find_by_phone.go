package usecase

import (
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/port"
	shared "github.com/TrueFlowDev/Backend/internal/shared/domain/value_object"
)

type FindUserByPhoneInput struct {
	Phone string
}

type FindUserByPhoneUsecase struct {
	userRepository port.UserRepository
}

func NewFindUserByPhoneUsecase(
	userRepository port.UserRepository,
) *FindUserByPhoneUsecase {
	return &FindUserByPhoneUsecase{
		userRepository: userRepository,
	}
}

func (u *FindUserByPhoneUsecase) Execute(input FindUserByPhoneInput) (*entity.User, error) {
	userPhone, err := shared.NewPhone(input.Phone)
	if err != nil {
		return nil, err
	}

	user, err := u.userRepository.FindByPhone(userPhone)
	if err != nil {
		return nil, err
	}

	return user, nil
}
