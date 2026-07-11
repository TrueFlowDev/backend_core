package port

import "github.com/TrueFlowDev/Backend/internal/module/user/domain/valueobject"

type UserIDGenerator interface {
	Generate() valueobject.UserID
}
