package http

import (
	"net/http"

	"github.com/ghosecorp/ghosecorp-auth/auth-api/internal/repository"
	"github.com/ghosecorp/ghosecorp-auth/auth-api/internal/security"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(sessionRepo *repository.SessionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(security.SessionCookieName)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		user, err := sessionRepo.FindUserBySessionTokenHash(
			c.Request.Context(),
			security.HashToken(token),
		)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
