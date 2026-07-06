package entity

import (
	"time"

	"github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"
)

type User struct {
	id       value_object.UserID
	phone    value_object.Phone
	password *value_object.HashedPassword

	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

type RestoreUserParams struct {
	ID       value_object.UserID
	Phone    value_object.Phone
	Password *value_object.HashedPassword

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewUser(
	id value_object.UserID,
	phone value_object.Phone,
) (*User, error) {
	now := time.Now().UTC()
	return &User{
		id:        id,
		phone:     phone,
		createdAt: now,
		updatedAt: now,
	}, nil
}

func RestoreUser(
	params RestoreUserParams,
) *User {
	return &User{
		id:        params.ID,
		phone:     params.Phone,
		password:  params.Password,
		createdAt: params.CreatedAt,
		updatedAt: params.UpdatedAt,
		deletedAt: params.DeletedAt,
	}
}

// <-- Getters -->

func (u *User) ID() value_object.UserID                { return u.id }
func (u *User) Phone() value_object.Phone              { return u.phone }
func (u *User) Password() *value_object.HashedPassword { return u.password }
func (u *User) CreatedAt() time.Time                   { return u.createdAt }
func (u *User) UpdatedAt() time.Time                   { return u.updatedAt }
func (u *User) DeletedAt() *time.Time                  { return u.deletedAt }

// <-- Setters -->

func (u *User) UpdatePassword(newPassword *value_object.HashedPassword) {
	u.password = newPassword
	u.touch()
}

// <-- Helpers -->

func (u *User) touch() {
	u.updatedAt = time.Now().UTC()
}
