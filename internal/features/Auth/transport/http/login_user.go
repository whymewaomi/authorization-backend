package auth_transport_http

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/whymewaomi/authorization-backend/internal/core/domain"
	core_dto "github.com/whymewaomi/authorization-backend/internal/core/dto"
	core_errors "github.com/whymewaomi/authorization-backend/internal/core/errors"
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
		c.JSON(http.StatusBadRequest, core_errors.ErrorsMessage{
			StatusCode: core_errors.ErrBadRequest.Error(),
			Details:    err.Error(),
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
		errors := core_errors.ErrorsResponse(err)
		c.JSON(errors.Status, errors)
		return
	}

	refreshToken := uuid.NewString()

	c.SetCookie("refresh_token", refreshToken, 9999999, "/", "", isProd, true)

	if err := h.authService.SaveRefreshToken(ctx, refreshToken, jwt.ID); err != nil {
		errors := core_errors.ErrorsResponse(err)
		c.JSON(errors.Status, errors)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": jwt.AccessToken,
	})
}
