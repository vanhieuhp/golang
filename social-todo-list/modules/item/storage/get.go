package storage

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

func (sql *sqlStorage) GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoIem, error) {

	var data model.TodoIem

	if err := sql.db.Where(cond).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrDB(err)
		}
		return nil, err
	}

	return &data, nil
}
