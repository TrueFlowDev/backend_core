package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/valueobject"
)

type ListPermissionsOutput struct {
	Category    string
	Permissions []PermissionOutput
}

type PermissionOutput struct {
	Value string
	Title string
}

type ListPermissionsUseCase struct{}

func NewListPermissionsUseCase() *ListPermissionsUseCase {
	return &ListPermissionsUseCase{}
}

func (uc *ListPermissionsUseCase) Execute(ctx context.Context) []ListPermissionsOutput {
	perms := valueobject.All()

	groups := make(map[string][]PermissionOutput)
	var order []string

	for _, p := range perms {
		cat := p.Category().String()
		if _, exists := groups[cat]; !exists {
			order = append(order, cat)
		}
		groups[cat] = append(groups[cat], PermissionOutput{
			Value: p.Value(),
			Title: p.Title(),
		})
	}

	result := make([]ListPermissionsOutput, 0, len(order))
	for _, cat := range order {
		result = append(result, ListPermissionsOutput{
			Category:    cat,
			Permissions: groups[cat],
		})
	}
	return result
}
