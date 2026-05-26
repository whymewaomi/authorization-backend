package auth_service

import (
	"context"
	"time"
)



func (s *AuthService) SaveRefreshToken(
	ctx context.Context,
	refreshToken string,
	userID interface{},
) error {
	return s.authStorage.Set(ctx, refreshToken, userID, 720 * time.Hour)
}
