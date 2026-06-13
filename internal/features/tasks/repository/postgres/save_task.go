package tasks_postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
	core_errors "github.com/Lesnekkk/golang-todo-app/internal/core/errors"
)

func (r *TasksPostgresRepository) Save(ctx context.Context, task domain.Task) (domain.Task, error) {
	query := `
	INSERT INTO todoapp.tasks (id, version, title, description, completed, created_at, completed_at, author_user_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id, version, title, description, completed, created_at, completed_at, author_user_id;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		task.ID,
		task.Version,
		task.Title,
		task.Description,
		task.Completed,
		task.CreatedAt,
		task.CompletedAt,
		task.AuthorUserID,
	)

	var saved domain.Task
	if err := row.Scan(
		&saved.ID,
		&saved.Version,
		&saved.Title,
		&saved.Description,
		&saved.Completed,
		&saved.CreatedAt,
		&saved.CompletedAt,
		&saved.AuthorUserID,
	); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return domain.Task{}, fmt.Errorf("user with id='%s': %w", task.AuthorUserID, core_errors.ErrNotFound)
		}
		return domain.Task{}, fmt.Errorf("save task: %w", err)
	}

	return saved, nil
}
