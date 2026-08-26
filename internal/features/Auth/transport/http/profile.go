package auth_transport_http

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	core_dto "github.com/whymewaomi/authorization-backend/internal/core/dto"
	core_errors "github.com/whymewaomi/authorization-backend/internal/core/errors"
)

// ProfileUserAPI godoc
// @Summary Get user profile
// @Description Returns authorized user profile
// @Tags auth
// @Security BearerAuth
// @Produce json
// @Success 200 {object} auth_dto.ProfileUserResponse
// @Failure 401 {object} auth_dto.ErrorResponse
// @Failure 500 {object} auth_dto.ErrorResponse
// @Router /api/v1/profile [get]
func (h *HTTPAuth) ProfileUserAPI(c *gin.Context) {
	val, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusBadRequest, core_errors.ErrorsMessage{
			StatusCode: core_errors.ErrBadRequest.Error(),
			Details:    "userID not found",
		})
		return
	}

	userID, ok := val.(int)
	if !ok {
		c.JSON(http.StatusBadRequest, core_errors.ErrorsMessage{
			StatusCode: core_errors.ErrBadRequest.Error(),
			Details:    "invalid userID",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	user, err := h.authService.ProfileUser(ctx, userID)
	if err != nil {
		errors := core_errors.ErrorsResponse(err)
		c.JSON(errors.Status, errors)
		return
	}

	c.JSON(http.StatusOK, core_dto.ProfileUserResponse{
		ID:         user.ID,
		Username:   user.Username,
		Email:      user.Email,
		RegisterAt: user.RegisterAt,
	})
}
