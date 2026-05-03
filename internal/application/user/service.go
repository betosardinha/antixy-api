package user

import (
	"context"

	"github.com/betosardinha/antixy-api/internal/domain/user"
)

type Service struct {
	repo user.Repository
}

func NewService(repo user.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(ctx context.Context, user user.User) (user.User, error) {
	return s.repo.Create(ctx, user)
}

func (s *Service) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	return s.repo.FindByEmail(ctx, email)
}
