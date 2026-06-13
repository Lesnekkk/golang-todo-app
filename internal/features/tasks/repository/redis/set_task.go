package tasks_redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
)

func (r *TasksRedisRepository) SetTask(ctx context.Context, task domain.Task) error {
	key := fmt.Sprintf("task:%s", task.ID)

	data, err := json.Marshal(task)
	if err != nil {
		return fmt.Errorf("marshal task: %w", err)
	}

	return r.client.Set(ctx, key, data, 5*time.Minute).Err()
}
