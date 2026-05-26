package auth_transport_http

import (
	auth_dto "auth/internal/features/Auth/transport/http/dto"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *HTTPAuth) LogoutUserAPI(c *gin.Context) {
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

	if err := h.authService.LogoutUser(ctx, cookie.Value); err != nil {
		c.JSON(http.StatusInternalServerError, auth_dto.ErrorResponse{
			Status: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.SetCookie("refresh_token", "", -1, "/", "", false, true)

  c.JSON(http.StatusOK, auth_dto.LogoutDtoReposnse{
		Status: "you have successfully logged out",
	})
}