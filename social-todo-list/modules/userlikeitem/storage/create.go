package storage

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/userlikeitem/model"
)

func (s *sqlStorage) Create(ctx context.Context, data *model.Like) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
