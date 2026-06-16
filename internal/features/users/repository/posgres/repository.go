package users_posgres_repository

import core_postgres_pool "github.com/Zakhar4uk/golang-app/internal/core/repository/postgres/pool"

type UserRepository struct {
	pool core_postgres_pool.Pool
}

func NewUserRepository(pool core_postgres_pool.Pool) *UserRepository {
	return &UserRepository{
		pool: pool,
	}
} 
