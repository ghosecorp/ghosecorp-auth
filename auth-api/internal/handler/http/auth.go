package http

import (
	"errors"
	"net/http"
	"strings"

	"github.com/ghosecorp/ghosecorp-auth/auth-api/internal/security"
	"github.com/ghosecorp/ghosecorp-auth/auth-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUseCase  *usecase.AuthUsecase
	cookieSecure bool
}

func NewAuthHandler(authUseCase *usecase.AuthUsecase, cookieSecure bool) *AuthHandler {
	return &AuthHandler{
		authUseCase:  authUseCase,
		cookieSecure: cookieSecure,
	}
}

type authRequest struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var req authRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body request"})
		return
	}

	if strings.TrimSpace(req.Email) == "" || len(req.Password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password of at least 8 characters are required"})
		return
	}

	user, err := h.authUseCase.Signup(c.Request.Context(), req.Email, req.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req authRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body request"})
		return
	}

	user, token, err := h.authUseCase.Login(c.Request.Context(), req.Email, req.Password)

	if err != nil {
		if errors.Is(err, usecase.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": "could not login"})
		return
	}

	security.SetSessionCookie(c, token, h.cookieSecure)
	c.JSON(http.StatusOK, gin.H{"user": user})
}
