package users_redis

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

func (r *UsersRedisRepository) GetByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	key := fmt.Sprintf("user:%s", id)

	data, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return domain.User{}, core_errors.ErrNotFound
		}
		return domain.User{}, fmt.Errorf("redis get user: %w", err)
	}

	var user domain.User
	if err := json.Unmarshal(data, &user); err != nil {
		return domain.User{}, fmt.Errorf("unmarshal user: %w", err)
	}

	return user, nil
}
