package auth_repository

import core_postgresql "github.com/whymewaomi/authorization-backend/internal/core/repository/postgresql"

type AuthRepository struct {
	pool core_postgresql.Pool
}

func NewAuthRepository(
	pool core_postgresql.Pool,
) *AuthRepository {
	return &AuthRepository{
		pool: pool,
	}
}
