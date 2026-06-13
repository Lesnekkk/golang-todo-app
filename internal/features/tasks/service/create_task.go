package tasks_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
)

func (s *TasksService) CreateTask(ctx context.Context, title string, description *string, authorUserID uuid.UUID) (domain.Task, error) {
	task := domain.CreateTask(title, description, authorUserID)

	if err := task.Validate(); err != nil {
		return domain.Task{}, fmt.Errorf("validate task: %w", err)
	}

	saved, err := s.repo.Save(ctx, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("save task: %w", err)
	}

	return saved, nil
}
