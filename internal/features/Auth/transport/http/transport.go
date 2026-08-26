package auth_transport_http

import (
	"context"

	"github.com/whymewaomi/authorization-backend/internal/core/domain"
	core_middleware "github.com/whymewaomi/authorization-backend/internal/core/middleware"

	"github.com/gin-gonic/gin"
)

type HTTPAuth struct {
	authService AuthService
	engine      *gin.Engine
}

type AuthService interface {
	SaveRefreshToken(
		ctx context.Context,
		refreshToken string,
		userID interface{},
	) error
	RegisterUser(
		ctx context.Context,
		user *domain.User,
	) (*domain.UserToken, error)
	RefreshToken(
		ctx context.Context,
		refreshToken string,
	) (string, error)
	LoginUser(
		ctx context.Context,
		user *domain.User,
	) (*domain.UserToken, error)
	LogoutUser(
		ctx context.Context,
		refreshToken string,
	) error
	ProfileUser(
		ctx context.Context,
		userID int,
	) (*domain.User, error)
}

func NewAuth(
	authService AuthService,
	engine *gin.Engine,
) *HTTPAuth {
	return &HTTPAuth{
		authService: authService,
		engine:      engine,
	}
}

func (h *HTTPAuth) RegisterRouter() {
	api := h.engine.Group("/api/v1")

	auth := api.Group("/auth")
	auth.POST("/register", h.RegisterUserAPI)
	auth.POST("/login", h.LoginUserAPI)
	auth.POST("/logout", h.LogoutUserAPI)

	token := api.Group("/token")
	token.POST("/refresh", h.RefreshTokenAPI)

	profile := api.Group("/profile")

	profile.GET("", core_middleware.JWTCheck(), h.ProfileUserAPI)
}
