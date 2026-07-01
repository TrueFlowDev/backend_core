package port

import user "github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"

type UserIdGenerator interface {
	Generate() user.UserID
}
