package entity

import (
	"strings"
	"time"

	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/module/authorization/domain/valueobject"
)

const (
	MinTitleLength = 2
	MaxTitleLength = 100
)

var (
	ErrTitleTooShort = xerr.New(
		xerr.CodeBadRequest,
		xerr.WithMeta("title", xerr.ErrorReasonTooShort),
	)

	ErrTitleTooLong = xerr.New(
		xerr.CodeBadRequest,
		xerr.WithMeta("title", xerr.ErrorReasonTooLong),
	)

	ErrDuplicatePermission = xerr.New(
		xerr.CodeBadRequest,
		xerr.WithMeta("permissions", xerr.ErrorReasonInvalidFormat),
	)

	ErrEmptyPermissions = xerr.New(
		xerr.CodeBadRequest,
		xerr.WithMeta("permissions", xerr.ErrorReasonRequired),
	)
)

type Role struct {
	id             valueobject.RoleID
	organizationID valueobject.OrganizationID

	title       string
	permissions []valueobject.Permission

	isOwner bool

	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

type RestoreRoleParams struct {
	ID             valueobject.RoleID
	OrganizationID valueobject.OrganizationID

	Title       string
	Permissions []valueobject.Permission

	IsOwner bool

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewRole(
	id valueobject.RoleID,
	organizationID valueobject.OrganizationID,
	title string,
	permissions []valueobject.Permission,
) (*Role, error) {
	title, err := validateTitle(title)
	if err != nil {
		return nil, err
	}

	if err := validatePermissions(permissions); err != nil {
		return nil, err
	}

	now := time.Now().UTC()

	return &Role{
		id:             id,
		organizationID: organizationID,
		title:          title,
		permissions:    permissions,
		createdAt:      now,
		updatedAt:      now,
	}, nil
}

func NewOwnerRole(
	id valueobject.RoleID,
	organizationID valueobject.OrganizationID,
) *Role {
	now := time.Now().UTC()

	return &Role{
		id:             id,
		organizationID: organizationID,
		title:          "مالک",
		isOwner:        true,
		createdAt:      now,
		updatedAt:      now,
	}
}

func RestoreRole(params RestoreRoleParams) *Role {
	return &Role{
		id:             params.ID,
		organizationID: params.OrganizationID,
		title:          params.Title,
		permissions:    params.Permissions,
		isOwner:        params.IsOwner,
		createdAt:      params.CreatedAt,
		updatedAt:      params.UpdatedAt,
		deletedAt:      params.DeletedAt,
	}
}

// <-- Getters -->

func (r *Role) ID() valueobject.RoleID                     { return r.id }
func (r *Role) OrganizationID() valueobject.OrganizationID { return r.organizationID }
func (r *Role) Title() string                              { return r.title }
func (r *Role) Permissions() []valueobject.Permission      { return r.permissions }
func (r *Role) IsOwner() bool                              { return r.isOwner }
func (r *Role) CreatedAt() time.Time                       { return r.createdAt }
func (r *Role) UpdatedAt() time.Time                       { return r.updatedAt }
func (r *Role) DeletedAt() *time.Time                      { return r.deletedAt }

// <-- Helpers -->

func validateTitle(title string) (string, error) {
	title = strings.TrimSpace(title)

	length := len([]rune(title))

	if length < MinTitleLength {
		return "", ErrTitleTooShort
	}

	if length > MaxTitleLength {
		return "", ErrTitleTooLong
	}

	return title, nil
}

func validatePermissions(permissions []valueobject.Permission) error {
	if len(permissions) == 0 {
		return ErrEmptyPermissions
	}

	seen := make(map[string]struct{}, len(permissions))

	for _, p := range permissions {
		if _, exists := seen[p.Value()]; exists {
			return ErrDuplicatePermission
		}
		seen[p.Value()] = struct{}{}
	}

	return nil
}
