package usecase

import (
	"context"
	"errors"

	"github.com/ghosecorp/ghosecorp-auth/auth-api/internal/domain"
	"github.com/ghosecorp/ghosecorp-auth/auth-api/internal/repository"
	"github.com/ghosecorp/ghosecorp-auth/auth-api/internal/security"
)

var ErrInvalidSession = errors.New("invalid session")

type SessionUseCase struct {
	sessionRepo *repository.SessionRepository
}

func NewSessionUseCase(sessionRepo *repository.SessionRepository) *SessionUseCase {
	return &SessionUseCase{
		sessionRepo: sessionRepo,
	}
}

func (u *SessionUseCase) GetUserBySessionToken(ctx context.Context, token string) (domain.User, error) {
	if token == "" {
		return domain.User{}, ErrInvalidSession
	}

	tokenHash := security.HashToken(token)

	user, err := u.sessionRepo.FindUserBySessionTokenHash(ctx, tokenHash)

	if err != nil {
		return domain.User{}, ErrInvalidSession
	}

	return user, nil
}
