package server

import (
	"net/http"

	"github.com/ghosecorp/ghosecorp-auth/auth-mailer/internal/email"
)

func NewRouter(emailHandler *email.Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /sample", emailHandler.SendEmail)
	return mux
}
