package store

import (
	"context"
	"database/sql"
	"time"
)

type Comment struct {
	Id        int64     `json:"id"`
	Content   string    `json:"content"`
	PostId    int64     `json:"post_id"`
	UserId    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user"`
}

type CommentStore struct {
	db *sql.DB
}

func (commentStore *CommentStore) CreateComment(ctx context.Context, comment *Comment) error {
	query := `
		INSERT INTO comments (content, post_id, user_id) 
		VALUES ($1, $2, $3) 
		RETURNING id, created_at
	`

	ctx, cancel := context.WithTimeout(ctx, TimeoutDuration)
	defer cancel()

	err := commentStore.db.QueryRowContext(
		ctx,
		query,
		comment.Content,
		comment.PostId,
		comment.UserId,
	).Scan(&comment.Id, &comment.CreatedAt)

	if err != nil {
		return err
	}

	return nil

}

func (commentStore *CommentStore) GetCommentsByPostId(ctx context.Context, postId int64) ([]Comment, error) {
	query := `
		select c.id, c.post_id, c.user_id, c.content, c.created_at, users.username, users.id 
		from comments c
		inner join users on users.id = c.user_id
		inner join posts on posts.id = c.post_id
		where c.post_id = $1
		order by c.created_at desc
    `

	ctx, cancel := context.WithTimeout(ctx, TimeoutDuration)
	defer cancel()

	rows, err := commentStore.db.QueryContext(ctx, query, postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []Comment{}
	for rows.Next() {
		var comment Comment
		comment.User = User{}
		err = rows.Scan(&comment.Id, &comment.PostId, &comment.UserId, &comment.Content, &comment.CreatedAt, &comment.User.Username, &comment.Id)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
