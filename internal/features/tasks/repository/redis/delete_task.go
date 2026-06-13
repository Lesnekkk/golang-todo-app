package tasks_redis

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (r *TasksRedisRepository) DeleteTask(ctx context.Context, id uuid.UUID) error {
	key := fmt.Sprintf("task:%s", id)
	return r.client.Del(ctx, key).Err()
}
