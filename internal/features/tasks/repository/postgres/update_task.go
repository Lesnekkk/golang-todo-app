package tasks_postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
	core_errors "github.com/Lesnekkk/golang-todo-app/internal/core/errors"
)

func (r *TasksPostgresRepository) Update(ctx context.Context, task domain.Task) (domain.Task, error) {
	query := `
	UPDATE todoapp.tasks
	SET title=$1, description=$2, completed=$3, completed_at=$4, version=version+1
	WHERE id=$5 AND version=$6
	RETURNING id, version, title, description, completed, created_at, completed_at, author_user_id;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		task.Title,
		task.Description,
		task.Completed,
		task.CompletedAt,
		task.ID,
		task.Version,
	)

	var updated domain.Task
	if err := row.Scan(
		&updated.ID,
		&updated.Version,
		&updated.Title,
		&updated.Description,
		&updated.Completed,
		&updated.CreatedAt,
		&updated.CompletedAt,
		&updated.AuthorUserID,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Task{}, fmt.Errorf("task with id='%s' concurrently accessed: %w", task.ID, core_errors.ErrConflict)
		}
		return domain.Task{}, fmt.Errorf("update task: %w", err)
	}

	return updated, nil
}
