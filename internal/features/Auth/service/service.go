package auth_service

import (
	"context"
	"time"

	"github.com/whymewaomi/authorization-backend/internal/core/domain"
)

type AuthService struct {
	authRepository AuthRepository
	authStorage    AuthStorage
}

type AuthRepository interface {
	GetUserByUsername(
		ctx context.Context,
		username string,
	) (*domain.User, error)
	CreateUser(
		ctx context.Context,
		user *domain.User,
	) (int, error)
	GetUserByID(
		ctx context.Context,
		userID int,
	) (*domain.User, error)
}

type AuthStorage interface {
	Set(
		ctx context.Context,
		key string,
		val interface{},
		ttl time.Duration,
	) error
	Get(
		ctx context.Context,
		key string,
	) (string, error)
	Del(
		ctx context.Context,
		key string,
	) error
}

func NewAuthService(
	authRepository AuthRepository,
	authStorage AuthStorage,
) *AuthService {
	return &AuthService{
		authRepository: authRepository,
		authStorage:    authStorage,
	}
}
