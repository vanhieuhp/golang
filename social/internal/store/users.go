package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(ctx context.Context, user *User) error {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id, created_at`

	ctx, cancel := context.WithTimeout(ctx, TimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Email,
		user.Password,
	).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		return err
	}
	return nil
}

func (userStore *UsersStore) GetById(ctx context.Context, id int64) (*User, error) {
	query := `
		SELECT id, username, email, password, created_at
		FROM users WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, TimeoutDuration)
	defer cancel()

	user := &User{}

	err := userStore.db.QueryRowContext(
		ctx,
		query,
		id).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrorNotFound
		default:
			return nil, err
		}
	}

	return user, nil

}
