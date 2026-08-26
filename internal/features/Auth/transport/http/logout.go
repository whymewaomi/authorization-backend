package auth_transport_http

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	core_dto "github.com/whymewaomi/authorization-backend/internal/core/dto"
)

// LogoutUserAPI godoc
// @Summary Logout user
// @Description Logout user and invalidate refresh token
// @Tags auth
// @Produce json
// @Success 200 {object} auth_dto.LogoutDtoReposnse
// @Failure 404 {object} auth_dto.ErrorResponse
// @Failure 500 {object} auth_dto.ErrorResponse
// @Router /api/v1/auth/logout [post]
func (h *HTTPAuth) LogoutUserAPI(c *gin.Context) {
	cookie, err := c.Request.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusNotFound, core_dto.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "invalid refresh token",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err := h.authService.LogoutUser(ctx, cookie.Value); err != nil {
		c.JSON(http.StatusInternalServerError, core_dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.SetCookie(
		"refresh_token",
		"",
		-1,
		"/",
		"",
		isProd,
		true,
	)

	c.Status(200)
}
