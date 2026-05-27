package auth_repository

import (
	"context"

	"github.com/whymewaomi/authorization-backend/internal/core/domain"
)

func (p *AuthRepository) GetUserByUsername(
	ctx context.Context,
	username string,
) (*domain.User, error) {
	sql := `
	SELECT id, username, email, password_hash, register_at
	FROM auth.user_auth 
	WHERE username = $1
	`

	var user domain.User

	if err := p.pool.QueryRow(ctx, sql, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.RegisterAt,
	); err != nil {
		return &domain.User{}, err
	}

  return &user, nil
}

func (p *AuthRepository) GetUserByID(
	ctx context.Context,
	userID int,
) (*domain.User, error) {
	sql := `
	SELECT id, username, email, register_at
	FROM auth.user_auth 
	WHERE id = $1
	`

	var user domain.User

	if err := p.pool.QueryRow(ctx, sql, userID).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.RegisterAt,
	); err != nil {
		return &domain.User{}, err
	}

  return &user, nil
}