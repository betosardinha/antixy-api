package repositories

import (
	"context"

	"github.com/betosardinha/antixy-api/internal/adapters/database/mappers"
	"github.com/betosardinha/antixy-api/internal/adapters/database/models"
	"github.com/betosardinha/antixy-api/internal/domain/user"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, u user.User) (user.User, error) {
	model := mappers.ToModel(u)

	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return user.User{}, err
	}

	return mappers.ToDomain(model), nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	var model models.UserModel

	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&model).
		Error

	if err != nil {
		return nil, err
	}

	domain := mappers.ToDomain(model)
	return &domain, nil
}
