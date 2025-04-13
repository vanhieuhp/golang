package storage

import (
	"context"
	"gorm.io/gorm"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

func (sql *sqlStorage) UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TodoItemUpdate) error {

	if err := sql.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return err
	}

	return nil
}

func (sql *sqlStorage) IncreaseLikeCount(ctx context.Context, id int) error {
	db := sql.db

	if err := db.Table(model.TodoIem{}.TableName()).Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (sql *sqlStorage) DecreaseLikeCount(ctx context.Context, id int) error {
	db := sql.db

	if err := db.Table(model.TodoIem{}.TableName()).Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
