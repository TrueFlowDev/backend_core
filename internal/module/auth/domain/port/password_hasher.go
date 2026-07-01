package port

import user "github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"

type PasswordHasher interface {
	Hash(password string) (string, error)
	Validate(password string, hashedPassword user.HashedPassword) (bool, error)
}
