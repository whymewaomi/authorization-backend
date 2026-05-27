package auth_service

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
	core_jwt "github.com/whymewaomi/authorization-backend/internal/core/jwt"
)


func (s *AuthService) RefreshToken(
	ctx context.Context,
	refreshToken string,
) (string, error) {
	userIDStr, err := s.authStorage.Get(ctx, refreshToken)
	if err != nil && !errors.Is(err, redis.Nil) {
		return "", fmt.Errorf("token invalid: %w", err)
	}

	if errors.Is(err, redis.Nil) {
		return "", errors.New("token invalid")
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return "", fmt.Errorf("error atoi: %w", err)
	}

	return core_jwt.GenerateJWT(userID)
}