package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrorNotFound   = errors.New("Resource not found")
	TimeoutDuration = time.Second * 5
	ErrorConflict   = errors.New("Resource already exists")
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetById(context.Context, int64) (*Post, error)
		Delete(context.Context, int64) error
		Update(context.Context, *Post) error
	}
	Users interface {
		Create(context.Context, *User) error
		GetById(context.Context, int64) (*User, error)
	}
	Comments interface {
		CreateComment(context.Context, *Comment) error
		GetCommentsByPostId(context.Context, int64) ([]Comment, error)
	}
	Followers interface {
		Follow(ctx context.Context, followerId, userId int64) error
		Unfollow(ctx context.Context, followerId, userId int64) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:     &PostStore{db},
		Users:     &UsersStore{db},
		Comments:  &CommentStore{db},
		Followers: &FollowerStore{db},
	}
}
