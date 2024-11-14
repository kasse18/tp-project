package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"tp-project/pkg/logger"
)

type Middleware struct {
	logger logger.Logger
}

func InitMiddleware(logger *logger.Logger) Middleware {
	return Middleware{
		logger: *logger,
	}
}

func (m Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			m.logger.Error(c.Request.Context(), "Authorization header missing")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			m.logger.Error(c.Request.Context(), "Bearer token missing")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bearer token missing"})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				m.logger.Error(c.Request.Context(), "Unexpected signing method")
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			secretKey := []byte("secret")
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			m.logger.Error(c.Request.Context(), "Invalid token", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Next()
	}
}
