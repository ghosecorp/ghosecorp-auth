package repository

import (
	"context"
	"database/sql"

	"github.com/ghosecorp/ghosecorp-auth/auth-api/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

type NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) createUser(ctx context.context, tx *sql.Tx, email string) (domain.User, error) {
	var User domain.user

	query := `
		INSERT INTO users (email)
		VALUES ($1)
		RETURNING user_id, public_id::text, email, is_active, created_at
	`

	err := tx.QueryRowContext(ctx, query, email).Scan(
		&user.UserID,
		&user.PublicID,
		&user.Email,
		&user.IsActive,
		&user.CreatedAt,
	)

	return user, err
}

func (r *UserRepository) FindUserByEmail (ctx context.Context, email string) (domain.User, string, error) {
	var user domain.User
	var passwordHash string

	query := `
		SELECT u.user_id, u.public_id::text, u.email, u.is_active, u.created_at, c.password_hash
		FROM users u
		JOIN credentials c ON c.user_id = u.user_id
		WHERE u.email = $1
	`
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.UserID,
		&user.PublicID,
		&user.Email,
		&user.IsActive,
		&user.CreatedAt,
		&passwordHash,
	)

	return user, passwordHash, err
}
