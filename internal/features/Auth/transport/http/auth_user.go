package auth_transport_http

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/whymewaomi/authorization-backend/internal/core/domain"
	core_dto "github.com/whymewaomi/authorization-backend/internal/core/dto"
)

var (
	isProd = false
)

// RegisterUserAPI godoc
// @Summary Register new user
// @Description Create a new account and return JWT tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param input body auth_dto.RegisterUserDTO true "Register data"
// @Success 200 {object} auth_dto.RegisterUserResponse
// @Failure 400 {object} auth_dto.ErrorResponse
// @Failure 500 {object} auth_dto.ErrorResponse
// @Router /api/v1/auth/register [post]
func (h *HTTPAuth) RegisterUserAPI(c *gin.Context) {
	var authUser core_dto.RegisterUserDTO
	if err := c.ShouldBindJSON(&authUser); err != nil {
		c.JSON(http.StatusBadRequest, core_dto.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	userDomain := domain.NewUser(
		authUser.Username,
		&authUser.Email,
		authUser.Password,
	)

	jwt, err := h.authService.RegisterUser(ctx, userDomain)
	if err != nil {
		c.JSON(http.StatusBadRequest, core_dto.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	refreshToken := uuid.NewString()

	c.SetCookie("refresh_token", refreshToken, 9999999, "/", "", isProd, true)

	if err := h.authService.SaveRefreshToken(ctx, refreshToken, jwt.ID); err != nil {
		c.JSON(http.StatusInternalServerError, core_dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, core_dto.RegisterUserResponse{
		Username:    authUser.Username,
		Email:       authUser.Email,
		AccessToken: jwt.AccessToken,
	})
}
