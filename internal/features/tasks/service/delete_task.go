package tasks_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *TasksService) DeleteTask(ctx context.Context, id uuid.UUID) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("delete task: %w", err)
	}

	s.cache.DeleteTask(ctx, id)

	return nil
}
