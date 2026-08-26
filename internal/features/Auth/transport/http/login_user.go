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

// LoginUserAPI godoc
// @Summary Login user
// @Description Login with username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param input body auth_dto.LoginUserDto true "Login data"
// @Success 200 {object} gin.H
// @Failure 400 {object} auth_dto.ErrorResponse
// @Router /api/v1/auth/login [post]
func (h *HTTPAuth) LoginUserAPI(c *gin.Context) {
	var user core_dto.LoginUserDto
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, core_dto.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	userDomain := domain.NewUser(
		user.Username,
		nil,
		user.Password,
	)

	jwt, err := h.authService.LoginUser(ctx, userDomain)
	if err != nil {
		c.JSON(http.StatusUnauthorized, core_dto.ErrorResponse{
			Status:  http.StatusUnauthorized,
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

	c.JSON(http.StatusOK, gin.H{
		"access_token": jwt.AccessToken,
	})
}
