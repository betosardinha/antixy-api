package mappers

import (
	"github.com/betosardinha/antixy-api/internal/adapters/database/models"
	"github.com/betosardinha/antixy-api/internal/domain/user"
)

func ToDomain(m models.UserModel) user.User {
	return user.User{
		ID:             m.ID,
		Name:           m.Name,
		Email:          m.Email,
		PasswordDigest: m.PasswordDigest,
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
	}
}

func ToModel(u user.User) models.UserModel {
	return models.UserModel{
		BaseModel:      models.BaseModel{},
		Name:           u.Name,
		Email:          u.Email,
		PasswordDigest: u.PasswordDigest,
	}
}
