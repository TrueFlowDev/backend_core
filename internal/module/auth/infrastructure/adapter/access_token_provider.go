package adapter

import (
	"fmt"

	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/value_object"
	"github.com/TrueFlowDev/Backend/internal/platform/config"
	"github.com/golang-jwt/jwt/v5"
)

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
		return "", xerr.Wrap(err, port.ErrFailedToSignToken.Code())
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
				return nil, xerr.Wrap(fmt.Errorf("invalid method: %v", token.Method), port.InvalidToken.Code())
			}
			return p.secret, nil
		},
	)

	if err != nil {
		return value_object.AccessTokenClaims{}, xerr.Wrap(err, port.InvalidToken.Code())
	}

	c, ok := token.Claims.(*tokenClaims)
	if !ok || !token.Valid {
		return value_object.AccessTokenClaims{}, port.InvalidToken
	}

	return value_object.NewAccessTokenClaims(
		c.Subject,
		c.IssuedAt.Time,
		c.ExpiresAt.Time,
	), nil
}
