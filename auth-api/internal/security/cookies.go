package security

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const SessionCookieName = "gc_session"

func SetSessionCookie(c *gin.Context, token string, secure bool) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     SessionCookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   secure,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
	})
}
