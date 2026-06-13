package tasks_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
)

func (s *TasksService) GetTasks(ctx context.Context, userID *uuid.UUID, limit *int, offset *int) ([]domain.Task, error) {
	tasks, err := s.repo.GetMany(ctx, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get tasks: %w", err)
	}
	return tasks, nil
}
