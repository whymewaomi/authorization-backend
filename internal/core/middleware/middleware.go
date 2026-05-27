package core_middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	core_jwt "github.com/whymewaomi/authorization-backend/internal/core/jwt"
)


var (
	userID = "user_id"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowedOrigins := map[string]struct{}{
			"localhost:5050": {},
			"http://localhost:5050": {},
		}

		origin := c.GetHeader("Origin")

		if _, ok := allowedOrigins[origin]; ok {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		  c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		  c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		  c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		}

		if c.Request.Method == http.MethodOptions {
			c.Writer.WriteHeader(204)
			return
		}

		c.Next()
	}
}

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


		c.Set(userID, claims.UserID)

		c.Next()
	}
}

