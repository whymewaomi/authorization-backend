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



func (h *HTTPAuth) RegisterUserAPI(c *gin.Context) {
	var authUser auth_dto.RegisterUserDTO
	if err := c.ShouldBindJSON(&authUser); err != nil {
		c.JSON(http.StatusBadRequest, auth_dto.ErrorResponse{
			Status: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5 * time.Second)
	defer cancel()

	userDomain := NewRegisterUserFromDTO(authUser)
	jwt, err := h.authService.RegisterUser(ctx, userDomain)
	if err != nil {
		c.JSON(http.StatusBadRequest, auth_dto.ErrorResponse{
			Status: http.StatusBadRequest,
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
  

	c.JSON(200, auth_dto.RegisterUserResponse{
		Username: authUser.Username,
		Email: authUser.Email,
		AccessToken: jwt.AccessToken,
		RefreshToken: refreshToken,
	})
}

func NewRegisterUserFromDTO(user auth_dto.RegisterUserDTO) *domain.User {
	return domain.NewRegisterUser(
		user.Username,
		user.Email,
		user.Password,
		user.ConfirmPassword,
	)
}