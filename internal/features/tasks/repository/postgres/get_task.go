package tasks_postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
	core_errors "github.com/Lesnekkk/golang-todo-app/internal/core/errors"
)

func (r *TasksPostgresRepository) GetByID(ctx context.Context, id uuid.UUID) (domain.Task, error) {
	query := `
	SELECT id, version, title, description, completed, created_at, completed_at, author_user_id
	FROM todoapp.tasks
	WHERE id=$1;
	`

	row := r.pool.QueryRow(ctx, query, id)
	var task domain.Task
	if err := row.Scan(
		&task.ID,
		&task.Version,
		&task.Title,
		&task.Description,
		&task.Completed,
		&task.CreatedAt,
		&task.CompletedAt,
		&task.AuthorUserID,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Task{}, fmt.Errorf("task with id='%s': %w", id, core_errors.ErrNotFound)
		}
		return domain.Task{}, fmt.Errorf("get task: %w", err)
	}

	return task, nil
}
