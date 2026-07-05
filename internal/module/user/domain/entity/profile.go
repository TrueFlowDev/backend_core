package entity

import (
	"strings"
	"time"

	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"
)

const (
	MinNameLength = 2
	MaxNameLength = 50

	MaxHeadlineLength = 100
	MaxBioLength      = 1000
)

var (
	ErrFirstNameTooShort = xerr.New(xerr.CodeBadRequest, xerr.WithMeta("first_name", xerr.ErrorReasonTooShort))
	ErrFirstNameTooLong  = xerr.New(xerr.CodeBadRequest, xerr.WithMeta("first_name", xerr.ErrorReasonTooLong))
	ErrLastNameTooShort  = xerr.New(xerr.CodeBadRequest, xerr.WithMeta("last_name", xerr.ErrorReasonTooShort))
	ErrLastNameTooLong   = xerr.New(xerr.CodeBadRequest, xerr.WithMeta("last_name", xerr.ErrorReasonTooLong))

	ErrHeadlineTooLong = xerr.New(xerr.CodeBadRequest, xerr.WithMeta("headline", xerr.ErrorReasonTooLong))
	ErrBioTooLong      = xerr.New(xerr.CodeBadRequest, xerr.WithMeta("bio", xerr.ErrorReasonTooLong))
)

type Profile struct {
	id        value_object.ProfileID
	userID    value_object.UserID
	email     value_object.Email
	firstName string
	lastName  string
	headline  string
	bio       string

	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

type RestoreProfileParams struct {
	ID        value_object.ProfileID
	UserID    value_object.UserID
	Email     value_object.Email
	FirstName string
	LastName  string
	Headline  string
	Bio       string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewProfile(
	id value_object.ProfileID,
	userID value_object.UserID,
	email value_object.Email,
	firstName, lastName string,
) (*Profile, error) {
	firstName = strings.TrimSpace(firstName)
	lastName = strings.TrimSpace(lastName)

	firstNameLength := len([]rune(firstName))
	if firstNameLength < MinNameLength {
		return nil, ErrFirstNameTooShort
	}
	if firstNameLength > MaxNameLength {
		return nil, ErrFirstNameTooLong
	}

	lastNameLength := len([]rune(lastName))
	if lastNameLength < MinNameLength {
		return nil, ErrLastNameTooShort
	}
	if lastNameLength > MaxNameLength {
		return nil, ErrLastNameTooLong
	}

	now := time.Now().UTC()
	return &Profile{
		id:        id,
		userID:    userID,
		email:     email,
		firstName: firstName,
		lastName:  lastName,
		createdAt: now,
		updatedAt: now,
	}, nil
}

func RestoreProfile(
	params RestoreProfileParams,
) *Profile {
	return &Profile{
		id:        params.ID,
		userID:    params.UserID,
		email:     params.Email,
		firstName: params.FirstName,
		lastName:  params.LastName,
		headline:  params.Headline,
		bio:       params.Bio,
		createdAt: params.CreatedAt,
		updatedAt: params.UpdatedAt,
		deletedAt: params.DeletedAt,
	}
}

// <-- Getters -->

func (p *Profile) ID() value_object.ProfileID  { return p.id }
func (p *Profile) UserID() value_object.UserID { return p.userID }
func (p *Profile) Email() value_object.Email   { return p.email }
func (p *Profile) FirstName() string           { return p.firstName }
func (p *Profile) LastName() string            { return p.lastName }
func (p *Profile) Headline() string            { return p.headline }
func (p *Profile) Bio() string                 { return p.bio }
func (p *Profile) CreatedAt() time.Time        { return p.createdAt }
func (p *Profile) UpdatedAt() time.Time        { return p.updatedAt }
func (p *Profile) DeletedAt() *time.Time       { return p.deletedAt }

// <-- Setters -->

func (p *Profile) UpdateEmail(newEmail value_object.Email) {
	p.email = newEmail
	p.touch()
}
func (p *Profile) UpdateHeadline(newHeadline string) error {
	newHeadline = strings.TrimSpace(newHeadline)
	if len([]rune(newHeadline)) > MaxHeadlineLength {
		return ErrHeadlineTooLong
	}
	if p.headline == newHeadline {
		return nil
	}

	p.headline = newHeadline
	p.touch()
	return nil
}

func (p *Profile) UpdateBio(newBio string) error {
	newBio = strings.TrimSpace(newBio)
	if len([]rune(newBio)) > MaxBioLength {
		return ErrBioTooLong
	}
	if p.bio == newBio {
		return nil
	}

	p.bio = newBio
	p.touch()
	return nil
}

// <-- Helpers -->

func (p *Profile) touch() {
	p.updatedAt = time.Now().UTC()
}
