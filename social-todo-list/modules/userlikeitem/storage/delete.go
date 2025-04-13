package storage

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/userlikeitem/model"
)

func (sql *sqlStorage) Delete(ctx context.Context, userId, itemId int) error {

	if err := sql.db.Table(model.Like{}.TableName()).
		Where("user_id = ? and item_id = ?", userId, itemId).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
