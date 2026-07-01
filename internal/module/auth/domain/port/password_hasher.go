package port

import "github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"

type PasswordHasher interface {
	Hash(password string) (string, error)
	Validate(password string, hashedPassword value_object.HashedPassword) (bool, error)
}
