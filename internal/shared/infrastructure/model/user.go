package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       string
	Phone    string
	Password *string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (*User) TableName() string { return "users" }
