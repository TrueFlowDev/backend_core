package mapper

import (
	"time"

	"github.com/TrueFlowDev/Backend/internal/module/user/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"
	"github.com/TrueFlowDev/Backend/internal/module/user/infrastructure/model"
	"gorm.io/gorm"
)

func ProfileModelToEntity(m *model.UsersProfile) (*entity.Profile, error) {
	userID := value_object.NewUserID(m.UserID)

	email, err := value_object.NewEmail(m.Email)
	if err != nil {
		return nil, err
	}

	var deletedAt *time.Time
	if m.DeletedAt.Valid {
		deletedAt = &m.DeletedAt.Time
	}

	return entity.RestoreProfile(entity.RestoreProfileParams{
		UserID:    userID,
		Email:     email,
		FirstName: m.FirstName,
		LastName:  m.LastName,
		Headline:  m.Headline,
		Bio:       m.Bio,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: deletedAt,
	}), nil
}

func ProfileEntityToModel(e *entity.Profile) *model.UsersProfile {
	var deletedAt gorm.DeletedAt
	if e.DeletedAt() != nil {
		deletedAt = gorm.DeletedAt{
			Time:  *e.DeletedAt(),
			Valid: true,
		}
	}

	return &model.UsersProfile{
		UserID:    e.UserID().Value(),
		Email:     e.Email().Value(),
		FirstName: e.FirstName(),
		LastName:  e.LastName(),
		Headline:  e.Headline(),
		Bio:       e.Bio(),
		CreatedAt: e.CreatedAt(),
		UpdatedAt: e.UpdatedAt(),
		DeletedAt: deletedAt,
	}
}
