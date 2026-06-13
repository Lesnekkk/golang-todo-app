package users_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *UsersService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	s.cache.DeleteUser(ctx, id)

	return nil
}
