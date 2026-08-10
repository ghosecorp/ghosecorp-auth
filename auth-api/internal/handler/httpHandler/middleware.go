package httpHandler

import (
	"net/http"

	"github.com/ghosecorp/ghosecorp-auth/auth-api/internal/security"
	"github.com/ghosecorp/ghosecorp-auth/auth-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(sessionUsecase *usecase.SessionUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(security.SessionCookieName)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		user, err := sessionUsecase.GetUserBySessionToken(c.Request.Context(), token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
