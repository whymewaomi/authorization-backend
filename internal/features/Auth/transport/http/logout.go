package auth_transport_http

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	core_errors "github.com/whymewaomi/authorization-backend/internal/core/errors"
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
		c.JSON(http.StatusBadRequest, core_errors.ErrorsMessage{
			StatusCode: core_errors.ErrBadRequest.Error(),
			Details:    err.Error(),
		})
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err := h.authService.LogoutUser(ctx, cookie.Value); err != nil {
		errors := core_errors.ErrorsResponse(err)
		c.JSON(errors.Status, errors)
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
