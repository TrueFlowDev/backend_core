package adapter

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/module/authentication/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/authentication/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/authentication/domain/value_object"
	"github.com/redis/go-redis/v9"
)

var _ port.OTPStore = (*OTPStore)(nil)

type otpDTO struct {
	Code      string    `json:"code"`
	Attempts  uint8     `json:"attempts"`
	ExpiresAt time.Time `json:"expires_at"`
}

type OTPStore struct {
	client *redis.Client
}

func NewOTPStore(client *redis.Client) *OTPStore {
	return &OTPStore{client: client}
}

func (s *OTPStore) Set(
	ctx context.Context,
	key value_object.Phone,
	value entity.OTP,
) error {
	otp := otpDTO{
		Code:      value.Code(),
		Attempts:  value.Attempts(),
		ExpiresAt: value.ExpiresAt(),
	}

	data, err := json.Marshal(otp)
	if err != nil {
		return xerr.Wrap(
			err, port.ErrOTPStore.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticReason, "marshal_failed"),
		)
	}

	ttl := time.Until(value.ExpiresAt())
	if err := s.client.Set(ctx, key.Value(), data, ttl).Err(); err != nil {
		return xerr.Wrap(
			err, port.ErrOTPStore.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "otp_store_set"),
		)
	}

	return nil
}

func (s *OTPStore) Get(
	ctx context.Context,
	key value_object.Phone,
) (entity.OTP, error) {
	var dto otpDTO

	payload, err := s.client.Get(ctx, key.Value()).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return entity.OTP{}, port.ErrOTPNotFound
		}

		return entity.OTP{}, xerr.Wrap(
			err, port.ErrOTPStore.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "otp_store_get"),
		)
	}

	if err := json.Unmarshal(payload, &dto); err != nil {
		return entity.OTP{}, xerr.Wrap(
			err, port.ErrOTPStore.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "corrupted_payload"),
		)
	}

	otp := entity.RestoreOTP(dto.Code, dto.Attempts, dto.ExpiresAt)

	return otp, nil
}

func (s *OTPStore) Delete(
	ctx context.Context,
	key value_object.Phone,
) error {
	if err := s.client.Del(ctx, key.Value()).Err(); err != nil {
		return xerr.Wrap(
			err, port.ErrOTPStore.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "otp_store_delete"),
		)
	}

	return nil
}
