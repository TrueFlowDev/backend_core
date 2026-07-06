package mapper

import (
	"time"

	"github.com/TrueFlowDev/Backend/internal/module/user/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"
	"github.com/TrueFlowDev/Backend/internal/module/user/infrastructure/model"
	"gorm.io/gorm"
)

func UserModelToEntity(m *model.User) (*entity.User, error) {
	userID := value_object.NewUserID(m.ID)

	userPhone, err := value_object.NewPhone(m.Phone)
	if err != nil {
		return nil, err
	}

	var deletedAt *time.Time
	if m.DeletedAt.Valid {
		deletedAt = &m.DeletedAt.Time
	}

	params := entity.RestoreUserParams{
		ID:        userID,
		Phone:     userPhone,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: deletedAt,
	}

	if m.Password != nil {
		userPassword, err := value_object.NewHashedPassword(*m.Password)
		if err != nil {
			return nil, err
		}
		params.Password = &userPassword
	}

	user := entity.RestoreUser(params)

	return user, nil
}

func UserEntityToModel(e *entity.User) *model.User {
	var userPassword *string
	if password := e.Password(); password != nil {
		userPassword = new(password.Value())
	}

	var deletedAt gorm.DeletedAt
	if e.DeletedAt() != nil {
		deletedAt = gorm.DeletedAt{Time: *e.DeletedAt(), Valid: true}
	}

	return &model.User{
		ID:        e.ID().Value(),
		Phone:     e.Phone().Value(),
		Password:  userPassword,
		CreatedAt: e.CreatedAt(),
		UpdatedAt: e.UpdatedAt(),
		DeletedAt: deletedAt,
	}
}
