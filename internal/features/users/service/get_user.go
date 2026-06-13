package users_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
)

func (s *UsersService) GetUser(ctx context.Context, id uuid.UUID) (domain.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user: %w", err)
	}
	return user, nil
}
