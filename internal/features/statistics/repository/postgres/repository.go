package statistics_postgres

import "github.com/jackc/pgx/v5/pgxpool"

type StatisticsPostgresRepository struct {
	pool *pgxpool.Pool
}

func NewStatisticsPostgresRepository(pool *pgxpool.Pool) *StatisticsPostgresRepository {
	return &StatisticsPostgresRepository{pool: pool}
}
