package tasks_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
)

func (s *TasksService) GetTask(ctx context.Context, id uuid.UUID) (domain.Task, error) {
	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get task: %w", err)
	}
	return task, nil
}
