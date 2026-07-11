package usecase

import (
	"context"
	"time"

	"github.com/TrueFlowDev/Backend/internal/module/authentication/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/authentication/domain/valueobject"
	"github.com/TrueFlowDev/Backend/internal/platform/config"
	"go.uber.org/fx"
)

type LoginInput struct {
	Phone    string
	Password string
}

type LoginOutput struct {
	AccessToken string
}

type LoginUsecase struct {
	accessTokenProvider port.AccessTokenProvider
	passwordHasher      port.PasswordHasher
	accessTokenDuration time.Duration
	userFinderByPhone   port.UserFinderByPhone
}

type LoginParams struct {
	fx.In

	Config              *config.Config
	AccessTokenProvider port.AccessTokenProvider
	PasswordHasher      port.PasswordHasher
	UserFinderByPhone   port.UserFinderByPhone
}

func NewLoginUsecase(p LoginParams) *LoginUsecase {
	return &LoginUsecase{
		accessTokenProvider: p.AccessTokenProvider,
		passwordHasher:      p.PasswordHasher,
		accessTokenDuration: p.Config.JWT.AccessTTL,
		userFinderByPhone:   p.UserFinderByPhone,
	}
}

func (u *LoginUsecase) Execute(ctx context.Context, input LoginInput) (LoginOutput, error) {
	user, err := u.userFinderByPhone.FindByPhone(ctx, port.UserFinderByPhoneInput{Phone: input.Phone})
	if err != nil {
		return LoginOutput{}, err
	}

	if err := u.passwordHasher.Validate(input.Password, user.HashedPassword); err != nil {
		return LoginOutput{}, err
	}

	now := time.Now().UTC()
	expiresAt := now.Add(u.accessTokenDuration)
	tokenClaims := valueobject.NewAccessTokenClaims(user.ID, now, expiresAt)

	accessToken, err := u.accessTokenProvider.Generate(tokenClaims)
	if err != nil {
		return LoginOutput{}, err
	}

	return LoginOutput{AccessToken: accessToken}, nil
}
