package users_redis

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (r *UsersRedisRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	key := fmt.Sprintf("user:%s", id)
	return r.client.Del(ctx, key).Err()
}
