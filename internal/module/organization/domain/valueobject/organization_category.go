package valueobject

import (
	"github.com/Ali127Dev/xerr"
)

var (
	ErrInvalidOrganizationCategory = xerr.New(xerr.CodeBadRequest,
		xerr.WithMeta("organization_category", xerr.ErrorReasonInvalidFormat))
)

type OrganizationCategory struct {
	value string
}

var (
	OrganizationCategoryTechnology    = OrganizationCategory{"technology"}
	OrganizationCategoryFinance       = OrganizationCategory{"finance"}
	OrganizationCategoryRetail        = OrganizationCategory{"retail"}
	OrganizationCategoryManufacturing = OrganizationCategory{"manufacturing"}
	OrganizationCategoryOther         = OrganizationCategory{"other"}
)

var organizationCategories = buildOrganizationCategoryMap(
	OrganizationCategoryTechnology,
	OrganizationCategoryFinance,
	OrganizationCategoryRetail,
	OrganizationCategoryManufacturing,
	OrganizationCategoryOther,
)

func buildOrganizationCategoryMap(categories ...OrganizationCategory) map[string]OrganizationCategory {
	m := make(map[string]OrganizationCategory, len(categories))
	for _, c := range categories {
		m[c.value] = c
	}
	return m
}

func ParseOrganizationCategory(raw string) (OrganizationCategory, error) {
	c, ok := organizationCategories[raw]
	if !ok {
		return OrganizationCategory{}, ErrInvalidOrganizationCategory
	}
	return c, nil
}

func (c OrganizationCategory) Value() string {
	return c.value
}
