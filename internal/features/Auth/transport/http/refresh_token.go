package auth_transport_http

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	core_errors "github.com/whymewaomi/authorization-backend/internal/core/errors"
)

// RefreshTokenAPI godoc
// @Summary Refresh access token
// @Description Generate new access token using refresh token
// @Tags auth
// @Produce json
// @Success 200 {object} gin.H
// @Failure 401 {object} auth_dto.ErrorResponse
// @Failure 404 {object} auth_dto.ErrorResponse
// @Router /api/v1/token/refresh [post]
func (h *HTTPAuth) RefreshTokenAPI(c *gin.Context) {
	cookie, err := c.Request.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, core_errors.ErrorsMessage{
			StatusCode: core_errors.ErrBadRequest.Error(),
			Details:    "invalid refresh token",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	jwt, err := h.authService.RefreshToken(ctx, cookie.Value)
	if err != nil {
		errors := core_errors.ErrorsResponse(err)
		c.JSON(errors.Status, errors)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": jwt,
	})
}
