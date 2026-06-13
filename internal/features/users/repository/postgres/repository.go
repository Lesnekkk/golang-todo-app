package users_postgres

import "github.com/jackc/pgx/v5/pgxpool"

type UsersPostgresRepository struct {
	pool *pgxpool.Pool
}

func NewUserPostgresRepository(pool *pgxpool.Pool) *UsersPostgresRepository {
	return &UsersPostgresRepository{pool: pool}
}
