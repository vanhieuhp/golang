package storage

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

func (sql *sqlStorage) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {

	if err := sql.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
