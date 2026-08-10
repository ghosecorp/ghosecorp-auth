package main

import (
	"database/sql"
	"log"

	"github.com/ghosecorp/ghosecorp-auth/auth-api/internal/config"
	"github.com/ghosecorp/ghosecorp-auth/auth-api/internal/handler/httpHandler"
	"github.com/ghosecorp/ghosecorp-auth/auth-api/internal/repository"
	"github.com/ghosecorp/ghosecorp-auth/auth-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db, err := sql.Open("postgres", cfg.DatabaseURL)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(db)
	credentialRepo := repository.NewCredentialRepository(db)
	sessionRepo := repository.NewSessionRepository(db)

	authCase := usecase.NewAuthUsecase(db, userRepo, credentialRepo, sessionRepo)

	sessionUsecase := usecase.NewSessionUseCase(sessionRepo)

	authHandler := httpHandler.NewAuthHandler(authCase, cfg.CookieSecure)
	meHandler := httpHandler.NewMeHandler()
	authMiddleware := httpHandler.AuthMiddleware(sessionUsecase)

	router := gin.Default()

	httpHandler.RegisterRoutes(
		router,
		authHandler,
		meHandler,
		authMiddleware,
	)

	log.Println("auth-api running on :" + cfg.Port)

	if err := router.Run(":", cfg.Port); err != nil {
		log.Fatal(err)
	}
}
