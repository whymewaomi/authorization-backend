package auth_transport_http

import (
	auth_dto "auth/internal/features/Auth/transport/http/dto"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *HTTPAuth) RefreshTokenAPI(c *gin.Context) {
	cookie, err := c.Request.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusNotFound, auth_dto.ErrorResponse{
			Status: http.StatusNotFound,
			Message: "invalid refresh token",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5 * time.Second)
	defer cancel()

  jwt, err := h.authService.RefreshToken(ctx, cookie.Value)
	if err != nil {
		c.JSON(http.StatusUnauthorized, auth_dto.ErrorResponse{
			Status: http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, auth_dto.NewAccessToken{
		AccessToken: jwt,
	})
}