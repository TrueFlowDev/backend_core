package entity

import (
	"strings"
	"time"

	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/value_object"
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
	id       value_object.OrganizationID
	category value_object.OrganizationCategory
	name     string
	active   bool

	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

type RestoreOrganizationParams struct {
	ID       value_object.OrganizationID
	Category value_object.OrganizationCategory
	Name     string
	Active   bool

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewOrganization(
	id value_object.OrganizationID,
	name string,
	category value_object.OrganizationCategory,
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

func (o *Organization) ID() value_object.OrganizationID             { return o.id }
func (o *Organization) Category() value_object.OrganizationCategory { return o.category }
func (o *Organization) Name() string                                { return o.name }
func (o *Organization) Active() bool                                { return o.active }
func (o *Organization) CreatedAt() time.Time                        { return o.createdAt }
func (o *Organization) UpdatedAt() time.Time                        { return o.updatedAt }
func (o *Organization) DeletedAt() *time.Time                       { return o.deletedAt }

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
