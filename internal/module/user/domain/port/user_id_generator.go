package port

import "github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"

type UserIDGenerator interface {
	Generate() value_object.UserID
}
