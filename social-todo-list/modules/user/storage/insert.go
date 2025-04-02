package storage

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/user/model"
)

func (sql *sqlStorage) CreateUser(ctx context.Context, data *model.UserCreate) error {
	db := sql.db.Begin()

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
