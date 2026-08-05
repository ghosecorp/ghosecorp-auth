package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/ghosecorp/ghosecorp-auth/auth-api/internal/domain"
)

type SessionRepository struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

func (r *SessionRepository) CreateSession(ctx context.Context, userID int64, tokenHash string, expiresAt time.Time) error {
	query := `
		INSERT INTO sessions (user_id, refresh_token_hash, expires_at)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.ExecContext(ctx, query, userID, tokenHash, expiresAt)
	return err
}

func (r *SessionRepository) FindUserBySessionTokenHash(ctx context.Context, tokenHash string) (domain.User, error) {
	var user domain.User

	query := `
		SELECT u.user_id, u.public_id::text, u.email, u.is_active, u.created_at
		FROM sessions s
		JOIN users u ON u.user_id = s.user_id
		WHERE s.refresh_token_hash = $1
		  AND s.expires_at > NOW()
		  AND u.is_active = TRUE
	`

	err := r.db.QueryRowContext(ctx, query, tokenHash).Scan(
		&user.UserID,
		&user.PublicID,
		&user.Email,
		&user.IsActive,
		&user.CreatedAt,
	)

	return user, err
}
