package statistics_postgres_repository

import (
	core_postgres_pool "github.com/nurkenspashev92/todoapp/internal/core/repository/postgres/pool"
	statistics_repository "github.com/nurkenspashev92/todoapp/internal/features/statistics/repository"
)

var _ statistics_repository.StatisticsRepositoryInterface = (*StatisticsRepository)(nil)

type StatisticsRepository struct {
	pool core_postgres_pool.Pool
}

func NewStatisticsRepository(pool core_postgres_pool.Pool) *StatisticsRepository {
	return &StatisticsRepository{pool: pool}
}
