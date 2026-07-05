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

	var createdAt, updatedAt time.Time
	if m.CreatedAt != nil {
		createdAt = *m.CreatedAt
	}
	if m.UpdatedAt != nil {
		updatedAt = *m.UpdatedAt
	}

	var deletedAt *time.Time
	if m.DeletedAt.Valid {
		deletedAt = &m.DeletedAt.Time
	}

	user := entity.RestoreUser(userID, userPhone, createdAt, updatedAt, deletedAt)

	if m.Password != nil {
		userPassword, err := value_object.NewHashedPassword(*m.Password)
		if err != nil {
			return nil, err
		}
		user.ChangePassword(userPassword)
	}

	return user, nil
}

func UserEntityToModel(e *entity.User) *model.User {
	var userPassword *string
	if pw := e.Password().Value(); pw != "" {
		userPassword = &pw
	}

	var deletedAt gorm.DeletedAt
	if e.DeletedAt() != nil {
		deletedAt = gorm.DeletedAt{Time: *e.DeletedAt(), Valid: true}
	}

	return &model.User{
		ID:        e.ID().Value(),
		Phone:     e.Phone().Value(),
		Password:  userPassword,
		CreatedAt: new(e.CreatedAt()),
		UpdatedAt: new(e.UpdatedAt()),
		DeletedAt: deletedAt,
	}
}
