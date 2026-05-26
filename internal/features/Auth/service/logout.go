package auth_service

import "context"

func (s *AuthService) LogoutUser(
	ctx context.Context,
	refreshToken string, 
) error {
  return s.authStorage.Del(ctx, refreshToken)
}