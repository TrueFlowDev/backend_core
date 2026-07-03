package port

import "context"

type UserFinderByPhoneInput struct {
	Phone string
}

type UserFinderByPhoneOutput struct {
	ID             string
	HashedPassword string
}

type UserFinderByPhone interface {
	FindByPhone(ctx context.Context, input UserFinderByPhoneInput) (UserFinderByPhoneOutput, error)
}
