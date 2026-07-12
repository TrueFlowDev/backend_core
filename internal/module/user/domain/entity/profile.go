package entity

import (
	"strings"
	"time"

	"github.com/Ali127Dev/xerr"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/valueobject"
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
	userID    valueobject.UserID
	email     valueobject.Email
	firstName string
	lastName  string
	headline  string
	bio       string

	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

type RestoreProfileParams struct {
	UserID    valueobject.UserID
	Email     valueobject.Email
	FirstName string
	LastName  string
	Headline  string
	Bio       string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewProfile(
	userID valueobject.UserID,
	email valueobject.Email,
	firstName, lastName string,
) (*Profile, error) {
	firstName, err := validateFirstName(firstName)
	if err != nil {
		return nil, err
	}

	lastName, err = validateLastName(lastName)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()

	return &Profile{
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

func (p *Profile) UserID() valueobject.UserID { return p.userID }
func (p *Profile) Email() valueobject.Email   { return p.email }
func (p *Profile) FirstName() string          { return p.firstName }
func (p *Profile) LastName() string           { return p.lastName }
func (p *Profile) Headline() string           { return p.headline }
func (p *Profile) Bio() string                { return p.bio }
func (p *Profile) CreatedAt() time.Time       { return p.createdAt }
func (p *Profile) UpdatedAt() time.Time       { return p.updatedAt }
func (p *Profile) DeletedAt() *time.Time      { return p.deletedAt }

// <-- Setters -->

func (p *Profile) UpdateEmail(newEmail valueobject.Email) {
	if p.email == newEmail {
		return
	}

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
func (p *Profile) UpdateFirstName(newFirstName string) error {
	firstName, err := validateFirstName(newFirstName)
	if err != nil {
		return err
	}

	if p.firstName == firstName {
		return nil
	}

	p.firstName = firstName
	p.touch()

	return nil
}
func (p *Profile) UpdateLastName(newLastName string) error {
	lastName, err := validateLastName(newLastName)
	if err != nil {
		return err
	}

	if p.lastName == lastName {
		return nil
	}

	p.lastName = lastName
	p.touch()

	return nil
}

// <-- Helpers -->

func (p *Profile) touch() {
	p.updatedAt = time.Now().UTC()
}
func validateFirstName(firstName string) (string, error) {
	firstName = strings.TrimSpace(firstName)

	length := len([]rune(firstName))
	if length < MinNameLength {
		return "", ErrFirstNameTooShort
	}
	if length > MaxNameLength {
		return "", ErrFirstNameTooLong
	}

	return firstName, nil
}
func validateLastName(lastName string) (string, error) {
	lastName = strings.TrimSpace(lastName)

	length := len([]rune(lastName))
	if length < MinNameLength {
		return "", ErrLastNameTooShort
	}
	if length > MaxNameLength {
		return "", ErrLastNameTooLong
	}

	return lastName, nil
}
