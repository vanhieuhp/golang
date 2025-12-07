package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
)

type Post struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	UserID    int64     `json:"user_id"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Comments  []Comment `json:"comments"`
	Version   int64     `json:"version"`
}

type PostStore struct {
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	query := `
		INSERT INTO posts (content, title, user_id, tags) 
		VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at, version;
	`

	ctx, cancel := context.WithTimeout(ctx, TimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserID,
		pq.Array(post.Tags),
	).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt, &post.Version)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostStore) GetById(ctx context.Context, id int64) (*Post, error) {
	query := `SELECT id, content, title, user_id, tags, created_at, updated_at, version
			FROM posts WHERE id = $1`
	post := Post{}

	ctx, cancel := context.WithTimeout(ctx, TimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(&post.ID, &post.Content, &post.Title, &post.UserID, pq.Array(&post.Tags), &post.CreatedAt, &post.UpdatedAt, &post.Version)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrorNotFound
		default:
			return nil, err
		}
	}

	return &post, nil
}

func (s *PostStore) Delete(ctx context.Context, postId int64) error {
	query := `DELETE FROM posts WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, TimeoutDuration)
	defer cancel()

	res, err := s.db.ExecContext(ctx, query, postId)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrorNotFound
	}

	return nil
}

func (s *PostStore) Update(ctx context.Context, post *Post) error {
	query := `
	update posts
	set title = $1, content = $2, version = version + 1
	where id = $3 and version = $4
	RETURNING version;
	`

	ctx, cancel := context.WithTimeout(ctx, TimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Title,
		post.Content,
		post.ID,
		post.Version,
	).Scan(&post.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrorNotFound
		default:
			return err
		}
	}

	return nil
}
