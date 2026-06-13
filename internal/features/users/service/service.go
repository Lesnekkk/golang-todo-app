package users_service

import (
	"context"

	"github.com/google/uuid"
	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User) (domain.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (domain.User, error)
	GetMany(ctx context.Context, limit *int, offset *int) ([]domain.User, error)
	Update(ctx context.Context, user domain.User) (domain.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type UsersService struct {
	repo UserRepository
}

func NewUsersService(repo UserRepository) *UsersService {
	return &UsersService{repo: repo}
}
