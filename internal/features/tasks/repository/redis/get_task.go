package tasks_redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
	core_errors "github.com/Lesnekkk/golang-todo-app/internal/core/errors"
)

func (r *TasksRedisRepository) GetByID(ctx context.Context, id uuid.UUID) (domain.Task, error) {
	key := fmt.Sprintf("task:%s", id)

	data, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return domain.Task{}, core_errors.ErrNotFound
		}
		return domain.Task{}, fmt.Errorf("redis get task: %w", err)
	}

	var task domain.Task
	if err := json.Unmarshal(data, &task); err != nil {
		return domain.Task{}, fmt.Errorf("unmarshal task: %w", err)
	}

	return task, nil
}
