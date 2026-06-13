package users_redis

import "github.com/redis/go-redis/v9"

type UsersRedisRepository struct {
	client *redis.Client
}

func NewUsersRedisRepository(client *redis.Client) *UsersRedisRepository {
	return &UsersRedisRepository{client: client}
}
