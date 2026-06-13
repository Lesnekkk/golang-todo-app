package users_postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
	core_errors "github.com/Lesnekkk/golang-todo-app/internal/core/errors"
)

func (r *UsersPostgresRepository) GetByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	query := `
	SELECT id, version, full_name, phone_number
	FROM todoapp.users
	WHERE id=$1;
	`

	row := r.pool.QueryRow(ctx, query, id)

	var user domain.User
	if err := row.Scan(&user.ID, &user.Version, &user.FullName, &user.PhoneNumber); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with id='%s': %w", id, core_errors.ErrNotFound)
		}
		return domain.User{}, fmt.Errorf("get user: %w", err)
	}

	return user, nil
}
