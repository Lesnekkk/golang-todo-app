package users_service

import (
	"context"
	"fmt"

	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
)

func (s *UsersService) CreateUser(ctx context.Context, fullName string, phoneNumber *string) (domain.User, error) {
	user := domain.CreateUser(fullName, phoneNumber)

	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate user: %w", err)
	}

	saved, err := s.repo.Save(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("save user: %w", err)
	}

	return saved, nil
}
