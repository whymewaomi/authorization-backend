package auth_repository

import (
	"context"

	"github.com/whymewaomi/authorization-backend/internal/core/domain"
)


func (p *AuthRepository) CreateUser(
	ctx context.Context,
	user *domain.User,
) (int, error) {
	sql := `
  INSERT INTO auth.user_auth (username, email, password_hash)
	VALUES ($1, $2, $3)
	RETURNING id
	`
 
	var userID int
	if err := p.pool.QueryRow(
		ctx, 
		sql, 
		user.Username, 
		user.Email, 
		user.Password,
	).Scan(&userID); err != nil {
		return 0, err
	}

	return userID, nil
}