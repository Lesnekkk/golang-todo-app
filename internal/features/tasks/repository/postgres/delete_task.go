package tasks_postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	core_errors "github.com/Lesnekkk/golang-todo-app/internal/core/errors"
)

func (r *TasksPostgresRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM todoapp.tasks WHERE id=$1;`

	cmdTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete task: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("task with id='%s': %w", id, core_errors.ErrNotFound)
	}

	return nil
}
