package main

import (
	"github.com/ghosecorp/ghosecorp-auth/auth-mailer/internal/config"
	"github.com/ghosecorp/ghosecorp-auth/auth-mailer/internal/email"
	"github.com/ghosecorp/ghosecorp-auth/auth-mailer/internal/server"
)

func main() {

	cfg := config.Load()

	emailService := email.NewService(
		cfg.APIKey,
		cfg.EmailURL,
	)

	emailHandler := email.NewHandler(emailService)

	router := server.NewRouter(emailHandler)


}
