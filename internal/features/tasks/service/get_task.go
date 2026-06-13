package tasks_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
	core_errors "github.com/Lesnekkk/golang-todo-app/internal/core/errors"
)

func (s *TasksService) GetTask(ctx context.Context, id uuid.UUID) (domain.Task, error) {
	task, err := s.cache.GetByID(ctx, id)
	if err == nil {
		return task, nil
	}

	if !errors.Is(err, core_errors.ErrNotFound) {
		return domain.Task{}, fmt.Errorf("cache get task: %w", err)
	}

	task, err = s.repo.GetByID(ctx, id)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get task: %w", err)
	}

	s.cache.SetTask(ctx, task)

	return task, nil
}
