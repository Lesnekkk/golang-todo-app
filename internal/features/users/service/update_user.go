package users_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
)

func (s *UsersService) PatchUser(ctx context.Context, id uuid.UUID, patch domain.UserPatch) (domain.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user: %w", err)
	}

	if err := user.ApplyPatch(patch); err != nil {
		return domain.User{}, fmt.Errorf("apply patch: %w", err)
	}

	updated, err := s.repo.Update(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("update user: %w", err)
	}

	return updated, nil
}
