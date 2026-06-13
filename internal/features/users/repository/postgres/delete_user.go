package users_postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	core_errors "github.com/Lesnekkk/golang-todo-app/internal/core/errors"
)

func (r *UsersPostgresRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM todoapp.users WHERE id=$1;`

	cmdTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("user with id='%s': %w", id, core_errors.ErrNotFound)
	}

	return nil
}
