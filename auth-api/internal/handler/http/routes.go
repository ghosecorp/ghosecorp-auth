package http

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	r *gin.Engine,
	authHandler *AuthHandler,
	meHandler *MeHandler,
	authMiddleware gin.HandlerFunc,
) {
	v1 := r.Group("/v1")
	v1.POST("/auth/signup", authHandler.SignUp)
	v1.POST("/auth/login", authHandler.Login)

	protected := v1.Group("")
	protected.Use(authMiddleware)
	protected.GET("/me", meHandler.Me)
}
