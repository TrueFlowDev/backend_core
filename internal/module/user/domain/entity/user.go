package entity

import (
	"time"

	"github.com/TrueFlowDev/Backend/internal/module/user/domain/valueobject"
)

type User struct {
	id       valueobject.UserID
	phone    valueobject.Phone
	password *valueobject.HashedPassword

	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

type RestoreUserParams struct {
	ID       valueobject.UserID
	Phone    valueobject.Phone
	Password *valueobject.HashedPassword

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewUser(
	id valueobject.UserID,
	phone valueobject.Phone,
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

func (u *User) ID() valueobject.UserID                { return u.id }
func (u *User) Phone() valueobject.Phone              { return u.phone }
func (u *User) Password() *valueobject.HashedPassword { return u.password }
func (u *User) CreatedAt() time.Time                  { return u.createdAt }
func (u *User) UpdatedAt() time.Time                  { return u.updatedAt }
func (u *User) DeletedAt() *time.Time                 { return u.deletedAt }

// <-- Setters -->

func (u *User) UpdatePassword(newPassword *valueobject.HashedPassword) {
	u.password = newPassword
	u.touch()
}

// <-- Helpers -->

func (u *User) touch() {
	u.updatedAt = time.Now().UTC()
}
