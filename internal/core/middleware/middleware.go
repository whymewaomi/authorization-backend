package core_middleware

import (
	core_jwt "auth/internal/core/jwt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)


var (
	userID = "user_id"
)


func JWTCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{
	     "message": "missing token",
       })
			return 
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.ParseWithClaims(tokenStr, &core_jwt.Claims{}, func(token *jwt.Token) (any, error) {
			return core_jwt.JwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{
	      "message": "invalid token",
       })
			return
		}

		claims := token.Claims.(*core_jwt.Claims)


		c.Set("user_id", claims.UserID)

		c.Next()
	}
}
