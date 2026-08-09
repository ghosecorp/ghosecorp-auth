package domain

import "time"

type User struct {
	UserID    int64     `json:"_"`
	PublicID  string    `json:"id"`
	Email     string    `json:"email"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}
