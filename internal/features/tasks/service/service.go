package tasks_service

import (
	"context"

	"github.com/google/uuid"
	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
)

type TaskRepository interface {
	Save(ctx context.Context, task domain.Task) (domain.Task, error)
	GetByID(ctx context.Context, id uuid.UUID) (domain.Task, error)
	GetMany(ctx context.Context, userID *uuid.UUID, limit *int, offset *int) ([]domain.Task, error)
	Update(ctx context.Context, task domain.Task) (domain.Task, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type TaskCache interface {
	GetByID(ctx context.Context, id uuid.UUID) (domain.Task, error)
	SetTask(ctx context.Context, task domain.Task) error
	DeleteTask(ctx context.Context, id uuid.UUID) error
}

type TasksService struct {
	repo  TaskRepository
	cache TaskCache
}

func NewTaskService(repo TaskRepository, cache TaskCache) *TasksService {
	return &TasksService{repo: repo, cache: cache}
}
