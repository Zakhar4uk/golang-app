package statistics_posgres_repository

import core_postgres_pool "github.com/Zakhar4uk/golang-app/internal/core/repository/postgres/pool"

type StatisticsRepository struct {
	pool core_postgres_pool.Pool
}

func NewStatisticsRepository(pool core_postgres_pool.Pool) *StatisticsRepository {
	return &StatisticsRepository{
		pool: pool,
	}
}
