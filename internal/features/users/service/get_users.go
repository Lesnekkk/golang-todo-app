package users_service

import (
	"context"
	"fmt"

	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
)

func (s *UsersService) GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error) {
	users, err := s.repo.GetMany(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get users: %w", err)
	}
	return users, nil
}
