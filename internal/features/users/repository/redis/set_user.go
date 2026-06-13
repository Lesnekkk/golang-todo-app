package users_redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
)

func (r *UsersRedisRepository) SetUser(ctx context.Context, user domain.User) error {
	key := fmt.Sprintf("user:%s", user.ID)

	data, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("marshal user: %w", err)
	}

	return r.client.Set(ctx, key, data, 5*time.Minute).Err()
}
