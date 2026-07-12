package port

import "context"

type OrganizationFinderOutput struct {
	ID       string
	Name     string
	Category string
}

type OrganizationFinder interface {
	FindByIDs(ctx context.Context, organizationIDs []string) ([]OrganizationFinderOutput, error)
}
