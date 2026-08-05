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
}
