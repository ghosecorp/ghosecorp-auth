package repository

import (
	"context"
	"database/sql"
)

type CredentialRepository struct {
	db *sql.DB
}

func NewCredentialRepository(db *sql.DB) *CredentialRepository {
	return &CredentialRepository{db: db}
}

func (r *CredentialRepository) CreateCredential(ctx context.Context, tx *sql.Tx, userID int64, passwordHash string) error {
	query := `
		INSERT INTO credentials (user_id, password_hash)
		VALUES ($1, $2)
	`

	_, err := tx.ExecContext(ctx, query, userID, passwordHash)
	return err
}
