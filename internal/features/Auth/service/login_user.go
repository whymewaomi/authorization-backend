package auth_service

import (
	"auth/internal/core/domain"
	core_jwt "auth/internal/core/jwt"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)


func (s *AuthService) ValidateUser(
	ctx context.Context,
	user *domain.User,
) (int, error) {
	userFromDB, err := s.authRepository.GetUserByUsername(ctx, user.Username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, fmt.Errorf("check username: %w", err)
	}

	if userFromDB == nil {
		return 0, fmt.Errorf("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(user.Password)); err != nil {
		return 0, errors.New("incorrect nickname or password")
	}


	return userFromDB.ID, nil
}

func (s *AuthService) LoginUser(
	ctx context.Context,
	user *domain.User,
) (*domain.UserToken, error) {
	userID, err := s.ValidateUser(ctx, user)
	if err != nil {
		return &domain.UserToken{}, errors.New("invalid credentials")
	}

	jwt, err := core_jwt.GenerateJWT(userID)
	if err != nil {
		return &domain.UserToken{}, fmt.Errorf("error generate jwt: %w", err)
	}

	return domain.NewUserToken(userID, jwt), nil
}