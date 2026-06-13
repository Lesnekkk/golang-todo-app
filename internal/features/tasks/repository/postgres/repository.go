package tasks_postgres

import "github.com/jackc/pgx/v5/pgxpool"

type TasksPostgresRepository struct {
	pool *pgxpool.Pool
}

func NewTasksPostgresRepository(pool *pgxpool.Pool) *TasksPostgresRepository {
	return &TasksPostgresRepository{pool: pool}
}
