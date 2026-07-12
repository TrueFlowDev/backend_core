package port

import "context"

type UserFinderOutput struct {
	ID        string
	Phone     string
	FirstName string
	LastName  string
}

type UserFinder interface {
	FindByPhone(ctx context.Context, phone string) (UserFinderOutput, error)
	FindByIDs(ctx context.Context, userIDs []string) ([]UserFinderOutput, error)
}
