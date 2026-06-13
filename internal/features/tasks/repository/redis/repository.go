package tasks_redis

import "github.com/redis/go-redis/v9"

type TasksRedisRepository struct {
	client *redis.Client
}

func NewTasksRedisRepository(client *redis.Client) *TasksRedisRepository {
	return &TasksRedisRepository{client: client}
}
