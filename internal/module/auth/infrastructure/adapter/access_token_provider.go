package adapter

import (
	"errors"
	"fmt"

	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/value_object"
	"github.com/TrueFlowDev/Backend/internal/platform/config"
	"github.com/golang-jwt/jwt/v5"
)

var _ port.AccessTokenProvider = (*JwtProvider)(nil)

type tokenClaims struct {
	jwt.RegisteredClaims
}

type JwtProvider struct {
	secret []byte
}

func NewJwtProvider(cfg *config.Config) *JwtProvider {
	return &JwtProvider{
		secret: []byte(cfg.JWT.AccessSecret),
	}
}

func (p *JwtProvider) Generate(
	claims value_object.AccessTokenClaims,
) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		tokenClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Subject:   claims.UserID(),
				IssuedAt:  jwt.NewNumericDate(claims.IssuedAt()),
				ExpiresAt: jwt.NewNumericDate(claims.ExpiresAt()),
			},
		},
	)

	signedToken, err := token.SignedString(p.secret)
	if err != nil {
		return "", xerr.Wrap(
			err,
			port.ErrFailedToSignToken.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "jwt_generate"),
		)
	}

	return signedToken, nil
}

func (p *JwtProvider) Verify(
	tokenString string,
) (value_object.AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&tokenClaims{},
		func(token *jwt.Token) (any, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("invalid method: %v", token.Method)
			}
			return p.secret, nil
		},
	)

	if err != nil {
		sentinel, reason := classifyJwtError(err)
		return value_object.AccessTokenClaims{}, xerr.Wrap(
			err,
			sentinel.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "jwt_verify"),
			xerr.WithDiagnostics(xerr.DiagnosticReason, reason),
		)
	}

	c, ok := token.Claims.(*tokenClaims)
	if !ok || !token.Valid {
		return value_object.AccessTokenClaims{}, xerr.Wrap(
			port.ErrInvalidToken,
			port.ErrInvalidToken.Code(),
			xerr.WithDiagnostics(xerr.DiagnosticOperation, "jwt_verify"),
			xerr.WithDiagnostics(xerr.DiagnosticReason, "invalid_claims_type"),
		)
	}

	return value_object.NewAccessTokenClaims(
		c.Subject,
		c.IssuedAt.Time,
		c.ExpiresAt.Time,
	), nil
}

func classifyJwtError(err error) (*xerr.Error, string) {
	switch {
	case errors.Is(err, jwt.ErrTokenExpired):
		return port.ErrTokenExpired, "token_expired"
	case errors.Is(err, jwt.ErrTokenNotValidYet):
		return port.ErrTokenExpired, "token_not_valid_yet"
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		return port.ErrInvalidToken, "invalid_signature"
	case errors.Is(err, jwt.ErrTokenMalformed):
		return port.ErrInvalidToken, "malformed"
	case errors.Is(err, jwt.ErrTokenInvalidClaims):
		return port.ErrInvalidToken, "invalid_claims"
	default:
		return port.ErrInvalidToken, "unknown"
	}
}
