package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/user/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"
)

type RegisterUserInput struct {
	Phone          string
	HashedPassword string
}

type RegisterUserOutput struct {
	ID string
}

type RegisterUserUsecase struct {
	userIDGenerator port.UserIDGenerator
	userRepository  port.UserRepository
}

func NewRegisterUserUsecase(
	userIDGenerator port.UserIDGenerator,
	userRepository port.UserRepository,
) *RegisterUserUsecase {
	return &RegisterUserUsecase{
		userIDGenerator: userIDGenerator,
		userRepository:  userRepository,
	}
}

func (u *RegisterUserUsecase) Execute(ctx context.Context, input RegisterUserInput) (RegisterUserOutput, error) {
	newUserPhone, err := value_object.NewPhone(input.Phone)
	if err != nil {
		return RegisterUserOutput{}, err
	}

	newUserHashedPassword, err := value_object.NewHashedPassword(input.HashedPassword)
	if err != nil {
		return RegisterUserOutput{}, err
	}

	newUserID := u.userIDGenerator.Generate()

	newUser, err := entity.NewUser(newUserID, newUserPhone)
	if err != nil {
		return RegisterUserOutput{}, err
	}
	newUser.UpdatePassword(&newUserHashedPassword)

	if err := u.userRepository.Create(ctx, newUser); err != nil {
		return RegisterUserOutput{}, err
	}

	return RegisterUserOutput{ID: newUserID.Value()}, nil
}
