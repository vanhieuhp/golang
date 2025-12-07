package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

type Follower struct {
	UserId     int64 `json:"user_id"`
	FollowerId int64 `json:"follower_id"`
	CreatedAt  int64 `json:"created_at"`
}

type FollowerStore struct {
	db *sql.DB
}

func (followerStore *FollowerStore) Follow(ctx context.Context, followerId, userId int64) error {
	query := `
		insert into followers(user_id, follower_id) values ($1, $2)
		`

	ctx, cancel := context.WithTimeout(ctx, TimeoutDuration)
	defer cancel()

	_, err := followerStore.db.ExecContext(ctx, query, userId, followerId)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code.Name() == "23505" {
			return ErrorConflict
		}
	}

	return nil
}

func (followerStore *FollowerStore) Unfollow(ctx context.Context, followerId, userId int64) error {
	query := `
		delete from followers where user_id = $1 and follower_id = $2
		`

	ctx, cancel := context.WithTimeout(ctx, TimeoutDuration)
	defer cancel()

	_, err := followerStore.db.ExecContext(ctx, query, userId, followerId)
	if err != nil {
		return err
	}

	return nil
}
