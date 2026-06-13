package users_postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
	core_errors "github.com/Lesnekkk/golang-todo-app/internal/core/errors"
)

func (r *UsersPostgresRepository) Update(ctx context.Context, user domain.User) (domain.User, error) {
	query := `
	UPDATE todoapp.users
	SET full_name=$1, phone_number=$2, version=version+1
	WHERE id=$3 AND version=$4
	RETURNING id, version, full_name, phone_number;
	`

	row := r.pool.QueryRow(ctx, query, user.FullName, user.PhoneNumber, user.ID, user.Version)

	var updated domain.User
	if err := row.Scan(&updated.ID, &updated.Version, &updated.FullName, &updated.PhoneNumber); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with id='%s' concurrently accessed: %w", user.ID, core_errors.ErrConflict)
		}
		return domain.User{}, fmt.Errorf("update user: %w", err)
	}

	return updated, nil
}
