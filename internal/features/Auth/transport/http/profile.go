package auth_transport_http

import (
	auth_dto "auth/internal/features/Auth/transport/http/dto"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *HTTPAuth) ProfileUserAPI(c *gin.Context) {
	val, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, auth_dto.ErrorResponse{
			Status: http.StatusUnauthorized,
			Message: "unauthorized",
		})
		return
	}

	userID, ok := val.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, auth_dto.ErrorResponse{
			Status: http.StatusUnauthorized,
			Message: "invalid user id",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5 * time.Second)
	defer cancel()

	user, err := h.authService.ProfileUser(ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, auth_dto.ErrorResponse{
			Status: http.StatusInternalServerError,
			Message: "failed to get profile",
		})
		return
	}

	c.JSON(http.StatusOK, auth_dto.ProfileUserResponse{
		ID: user.ID,
		Username: user.Username,
		Email: user.Email,
		RegisterAt: user.RegisterAt,
	})
}