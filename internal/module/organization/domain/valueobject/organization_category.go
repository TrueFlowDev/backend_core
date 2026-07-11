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

func NewOrganizationCategory(raw string) (OrganizationCategory, error) {
	switch raw {
	case OrganizationCategoryTechnology.value:
		return OrganizationCategoryTechnology, nil
	case OrganizationCategoryFinance.value:
		return OrganizationCategoryFinance, nil
	case OrganizationCategoryRetail.value:
		return OrganizationCategoryRetail, nil
	case OrganizationCategoryManufacturing.value:
		return OrganizationCategoryManufacturing, nil
	case OrganizationCategoryOther.value:
		return OrganizationCategoryOther, nil
	default:
		return OrganizationCategory{}, ErrInvalidOrganizationCategory
	}
}

func (c OrganizationCategory) Value() string {
	return c.value
}
