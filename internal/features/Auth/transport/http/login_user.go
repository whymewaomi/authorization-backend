package auth_transport_http

import (
	"auth/internal/core/domain"
	auth_dto "auth/internal/features/Auth/transport/http/dto"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *HTTPAuth) LoginUserAPI(c *gin.Context) {
	var user auth_dto.LoginUserDto
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, auth_dto.ErrorResponse{
			Status: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5 * time.Second)
	defer cancel()

	userDomain := NewUserPatchLogin(user)
	jwt, err := h.authService.LoginUser(ctx, userDomain)
	if err != nil {
		c.JSON(http.StatusUnauthorized, auth_dto.ErrorResponse{
			Status: http.StatusUnauthorized,
			Message: err.Error(),
		})
		return 
	}

	refreshToken := uuid.NewString()

	c.SetCookie("refresh_token", refreshToken, 9999999, "/", "", false, true)


	if err := h.authService.SaveRefreshToken(ctx, refreshToken, jwt.ID); err != nil {
		c.JSON(http.StatusInternalServerError, auth_dto.ErrorResponse{
			Status: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}


	c.JSON(http.StatusOK, auth_dto.LoginUserResponse{
		AccessToken: jwt.AccessToken,
		RefreshToken: refreshToken,
	})
}

func NewUserPatchLogin(req auth_dto.LoginUserDto) *domain.User {
	return domain.NewLoginUser(
		req.Username,
		req.Password,
	)
}