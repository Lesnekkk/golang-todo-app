package users_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
	core_errors "github.com/Lesnekkk/golang-todo-app/internal/core/errors"
)

func (s *UsersService) GetUser(ctx context.Context, id uuid.UUID) (domain.User, error) {
	// 1. смотрим в Redis
	user, err := s.cache.GetByID(ctx, id)
	if err == nil {
		return user, nil
	}

	// если ошибка не "не найдено" — что-то сломалось в Redis, идём в Postgres
	if !errors.Is(err, core_errors.ErrNotFound) {
		return domain.User{}, fmt.Errorf("cache get user: %w", err)
	}

	// 2. идём в Postgres
	user, err = s.repo.GetByID(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user: %w", err)
	}

	// 3. сохраняем в Redis на будущее (ошибку игнорируем — кэш не критичен)
	s.cache.SetUser(ctx, user)

	return user, nil
}
