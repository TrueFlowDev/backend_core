package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/TrueFlowDev/Backend/internal/module/authentication/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/authentication/domain/valueobject"
	"github.com/TrueFlowDev/Backend/internal/platform/config"
	"go.uber.org/fx"
)

type VerifyOTPAndRegisterInput struct {
	Phone    string
	Password string
	Code     string
}

type VerifyOTPAndRegisterOutput struct {
	AccessToken string
}

type VerifyOTPAndRegisterUsecase struct {
	otpStore            port.OTPStore
	userRegisterer      port.UserRegisterer
	accessTokenProvider port.AccessTokenProvider
	passwordHasher      port.PasswordHasher
	accessTokenDuration time.Duration
}

type VerifyOTPAndRegisterParams struct {
	fx.In

	Config              *config.Config
	OtpStore            port.OTPStore
	UserRegisterer      port.UserRegisterer
	AccessTokenProvider port.AccessTokenProvider
	PasswordHasher      port.PasswordHasher
}

func NewVerifyOTPAndRegisterUsecase(p VerifyOTPAndRegisterParams) *VerifyOTPAndRegisterUsecase {
	return &VerifyOTPAndRegisterUsecase{
		otpStore:            p.OtpStore,
		userRegisterer:      p.UserRegisterer,
		accessTokenProvider: p.AccessTokenProvider,
		passwordHasher:      p.PasswordHasher,
		accessTokenDuration: p.Config.JWT.AccessTTL,
	}
}

func (u *VerifyOTPAndRegisterUsecase) Execute(ctx context.Context, input VerifyOTPAndRegisterInput) (VerifyOTPAndRegisterOutput, error) {
	phone, err := valueobject.NewPhone(input.Phone)
	if err != nil {
		return VerifyOTPAndRegisterOutput{}, err
	}

	otp, err := u.otpStore.Get(ctx, phone)
	if err != nil {
		return VerifyOTPAndRegisterOutput{}, err
	}

	if err := otp.Verify(input.Code); err != nil {
		if setErr := u.otpStore.Set(ctx, phone, otp); setErr != nil {
			return VerifyOTPAndRegisterOutput{}, errors.Join(err, setErr)
		}

		return VerifyOTPAndRegisterOutput{}, err
	}

	_ = u.otpStore.Delete(ctx, phone)

	newUserHashedPassword, err := u.passwordHasher.Hash(input.Password)
	if err != nil {
		return VerifyOTPAndRegisterOutput{}, err
	}

	output, err := u.userRegisterer.Register(ctx, port.UserRegistererInput{
		Phone:          input.Phone,
		HashedPassword: newUserHashedPassword,
	})
	if err != nil {
		return VerifyOTPAndRegisterOutput{}, err
	}

	now := time.Now().UTC()
	expiresAt := now.Add(u.accessTokenDuration)
	tokenClaims := valueobject.NewAccessTokenClaims(output.ID, now, expiresAt)

	accessToken, err := u.accessTokenProvider.Generate(tokenClaims)
	if err != nil {
		return VerifyOTPAndRegisterOutput{}, err
	}

	return VerifyOTPAndRegisterOutput{AccessToken: accessToken}, nil
}
