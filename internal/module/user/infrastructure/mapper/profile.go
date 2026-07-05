package mapper

import (
	"time"

	"github.com/TrueFlowDev/Backend/internal/module/user/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"
	"github.com/TrueFlowDev/Backend/internal/module/user/infrastructure/model"
	"gorm.io/gorm"
)

func ProfileModelToEntity(m *model.UsersProfile) (*entity.Profile, error) {
	profileID := value_object.NewProfileID(m.ID)
	userID := value_object.NewUserID(m.UserID)

	email, err := value_object.NewEmail(m.Email)
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

	return entity.RestoreProfile(entity.RestoreProfileParams{
		ID:        profileID,
		UserID:    userID,
		Email:     email,
		FirstName: m.FirstName,
		LastName:  m.LastName,
		Headline:  m.Headline,
		Bio:       m.Bio,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
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
		ID:        e.ID().Value(),
		UserID:    e.UserID().Value(),
		Email:     e.Email().Value(),
		FirstName: e.FirstName(),
		LastName:  e.LastName(),
		Headline:  e.Headline(),
		Bio:       e.Bio(),
		CreatedAt: new(e.CreatedAt()),
		UpdatedAt: new(e.UpdatedAt()),
		DeletedAt: deletedAt,
	}
}
