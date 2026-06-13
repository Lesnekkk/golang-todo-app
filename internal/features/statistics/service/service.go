package statistics_service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
)

type StatisticsRepository interface {
	GetTasks(ctx context.Context, userID *uuid.UUID, from *time.Time, to *time.Time) ([]domain.Task, error)
}

type StatisticsService struct {
	repo StatisticsRepository
}

func NewStatisticsService(repo StatisticsRepository) *StatisticsService {
	return &StatisticsService{repo: repo}
}
