package auth_service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
	"github.com/whymewaomi/authorization-backend/internal/core/domain"
)

func (s *AuthService) ProfileUser(
	ctx context.Context,
	userID int,
) (*domain.User, error) {
	var userIDCache = fmt.Sprintf("user_id:%d", userID)

	userCache, err := s.authStorage.Get(ctx, userIDCache)
	if err == nil {
		var user domain.User

		if err := json.Unmarshal([]byte(userCache), &user); err != nil {
			return &domain.User{}, fmt.Errorf("failed unmarshal %w", err)
		}

		return &user, nil
	} 

	if !errors.Is(err, redis.Nil) {
		return nil, fmt.Errorf("redis get: %w", err)
	}

	user, err := s.authRepository.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
		return &domain.User{}, fmt.Errorf("user not found")
	}

		return &domain.User{}, fmt.Errorf("error: %w", err)
	}
  
	userMarshal, err := json.Marshal(user)
	if err != nil {
		return &domain.User{}, fmt.Errorf("error marshal: %w", err)
	}

	if err := s.authStorage.Set(ctx, userIDCache, userMarshal, 5 * time.Second); err != nil {
		return &domain.User{}, fmt.Errorf("failed set cache: %w", err)
	}

	return user, nil
}