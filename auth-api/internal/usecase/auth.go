package usecase

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/ghosecorp/ghosecorp-auth/auth-api/internal/domain"
	"github.com/ghosecorp/ghosecorp-auth/auth-api/internal/repository"
	"github.com/ghosecorp/ghosecorp-auth/auth-api/internal/security"
)

var ErrInvalidCredentials = errors.New("invalid email or password")

type AuthUsecase struct {
	db             *sql.DB
	userRepo       *repository.UserRepository
	credentialRepo *repository.CredentialRepository
	sessionRepo    *repository.SessionRepository
}

func NewAuthUsecase(
	db *sql.DB,
	userRepo *repository.UserRepository,
	credentialRepo *repository.CredentialRepository,
	sessionRepo *repository.SessionRepository,
) *AuthUsecase {
	return &AuthUsecase{
		db:             db,
		userRepo:       userRepo,
		credentialRepo: credentialRepo,
		sessionRepo:    sessionRepo,
	}
}

func (u *AuthUsecase) Signup(ctx context.Context, email string, password string) (domain.User, error) {
	email = strings.ToLower(strings.TrimSpace(email))

	passwordHash, err := security.HashPassword(password)
	if err != nil {
		return domain.User{}, err
	}

	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.User{}, err
	}
	defer tx.Rollback()

	user, err := u.userRepo.CreateUser(ctx, tx, email)
	if err != nil {
		return domain.User{}, err
	}

	err = u.credentialRepo.CreateCredential(ctx, tx, user.UserID, passwordHash)
	if err != nil {
		return domain.User{}, err
	}

	if err := tx.Commit(); err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (u *AuthUsecase) Login(ctx context.Context, email string, password string) (domain.User, string, error) {
	email = strings.ToLower(strings.TrimSpace(email))

	user, passwordHash, err := u.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		return domain.User{}, "", ErrInvalidCredentials
	}

	if !security.VerifyPassword(password, passwordHash) {
		return domain.User{}, "", ErrInvalidCredentials
	}

	token, err := security.NewSessionToken()
	if err != nil {
		return domain.User{}, "", err
	}

	tokenHash := security.HashToken(token)
	expiresAt := time.Now().Add(7 * 24 * time.Hour)

	err = u.sessionRepo.CreateSession(ctx, user.UserID, tokenHash, expiresAt)
	if err != nil {
		return domain.User{}, "", err
	}

	return user, token, nil
}
