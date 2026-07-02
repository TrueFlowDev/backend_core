package mapper

import (
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"
	"github.com/TrueFlowDev/Backend/internal/shared/infrastructure/model"
	"gorm.io/gorm"
)

func UserModelToEntity(model *model.User) (*entity.User, error) {
	userID := value_object.NewUserID(model.ID)
	userPhone, err := value_object.NewPhone(model.Phone)
	if err != nil {
		return nil, err
	}

	user := entity.RestoreUser(userID, userPhone, model.CreatedAt, model.UpdatedAt, &model.DeletedAt.Time)
	if model.Password != nil {
		userPassword, err := value_object.NewHashedPassword(*model.Password)
		if err != nil {
			return nil, err
		}
		user.ChangePassword(userPassword)
	}

	return user, nil
}

func UserEntityToModel(entity *entity.User) *model.User {
	var userPassword *string
	if entity.Password().Value() != "" {
		userPassword = new(entity.Password().Value())
	}

	deletedAt := gorm.DeletedAt{}
	if entity.DeletedAt() != nil {
		deletedAt = gorm.DeletedAt{
			Time:  *entity.DeletedAt(),
			Valid: true,
		}
	}

	return &model.User{
		ID:        entity.ID().Value(),
		Phone:     entity.Password().Value(),
		Password:  userPassword,
		CreatedAt: entity.CreatedAt(),
		UpdatedAt: entity.UpdatedAt(),
		DeletedAt: deletedAt,
	}
}
