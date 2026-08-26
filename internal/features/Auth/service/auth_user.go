package auth_service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	core_errors "github.com/whymewaomi/authorization-backend/internal/core/errors"
	core_jwt "github.com/whymewaomi/authorization-backend/internal/core/jwt"

	"github.com/whymewaomi/authorization-backend/internal/core/domain"

	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) HashPassword(
	password string,
) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(passwordHash), nil
}

func (s *AuthService) RegisterUser(
	ctx context.Context,
	user *domain.User,
) (*domain.UserToken, error) {
	if err := user.Validate(); err != nil {
		return &domain.UserToken{}, fmt.Errorf("validation error: %w", err)
	}

	existingUser, err := s.authRepository.GetUserByUsername(ctx, user.Username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return &domain.UserToken{}, core_errors.ErrValidation
	}
	if existingUser != nil {
		return &domain.UserToken{}, core_errors.ErrUsernameExists
	}

	passwordHash, err := s.HashPassword(user.Password)
	if err != nil {
		return &domain.UserToken{}, core_errors.ErrBadRequest
	}
	user.Password = passwordHash

	userID, err := s.authRepository.CreateUser(ctx, user)
	if err != nil {
		return &domain.UserToken{}, fmt.Errorf("failed create user: %w", err)
	}

	jwt, err := core_jwt.GenerateJWT(userID)
	if err != nil {
		return &domain.UserToken{}, fmt.Errorf("error create jwt: %w", err)
	}

	return domain.NewUserToken(userID, jwt), nil
}
