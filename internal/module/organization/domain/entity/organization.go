package entity

import (
	"strings"
	"time"

	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/valueobject"
)

const (
	MinOrganizationNameLength = 3
	MaxOrganizationNameLength = 100
)

var (
	ErrOrganizationNameTooShort = xerr.New(
		xerr.CodeBadRequest,
		xerr.WithMeta("organization_name", xerr.ErrorReasonTooShort),
	)

	ErrOrganizationNameTooLong = xerr.New(
		xerr.CodeBadRequest,
		xerr.WithMeta("organization_name", xerr.ErrorReasonTooLong),
	)
)

type Organization struct {
	id       valueobject.OrganizationID
	category valueobject.OrganizationCategory
	name     string
	active   bool

	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

type RestoreOrganizationParams struct {
	ID       valueobject.OrganizationID
	Category valueobject.OrganizationCategory
	Name     string
	Active   bool

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewOrganization(
	id valueobject.OrganizationID,
	name string,
	category valueobject.OrganizationCategory,
) (*Organization, error) {
	name, err := validateOrganizationName(name)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return &Organization{
		id:        id,
		name:      name,
		category:  category,
		active:    true,
		createdAt: now,
		updatedAt: now,
	}, nil
}

func RestoreOrganization(
	params RestoreOrganizationParams,
) *Organization {
	return &Organization{
		id:        params.ID,
		category:  params.Category,
		name:      params.Name,
		active:    params.Active,
		createdAt: params.CreatedAt,
		updatedAt: params.UpdatedAt,
		deletedAt: params.DeletedAt,
	}
}

// <-- Getters -->

func (o *Organization) ID() valueobject.OrganizationID             { return o.id }
func (o *Organization) Category() valueobject.OrganizationCategory { return o.category }
func (o *Organization) Name() string                               { return o.name }
func (o *Organization) Active() bool                               { return o.active }
func (o *Organization) CreatedAt() time.Time                       { return o.createdAt }
func (o *Organization) UpdatedAt() time.Time                       { return o.updatedAt }
func (o *Organization) DeletedAt() *time.Time                      { return o.deletedAt }

// <-- Setters -->

func (o *Organization) Update(name string, category valueobject.OrganizationCategory) error {
	validatedName, err := validateOrganizationName(name)
	if err != nil {
		return err
	}

	o.name = validatedName
	o.category = category
	o.updatedAt = time.Now().UTC()

	return nil
}

// <-- Helpers -->

func validateOrganizationName(name string) (string, error) {
	name = strings.TrimSpace(name)

	length := len([]rune(name))
	if length < MinOrganizationNameLength {
		return "", ErrOrganizationNameTooShort
	}
	if length > MaxOrganizationNameLength {
		return "", ErrOrganizationNameTooLong
	}

	return name, nil
}
