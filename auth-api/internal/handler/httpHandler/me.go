package httpHandler

import (
	"net/http"

	"github.com/ghosecorp/ghosecorp-auth/auth-api/internal/domain"
	"github.com/gin-gonic/gin"
)

type MeHandler struct{}

func NewMeHandler() *MeHandler {
	return &MeHandler{}
}

func (h *MeHandler) Me(c *gin.Context) {
	value, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	user, ok := value.(domain.User)

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user context"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
