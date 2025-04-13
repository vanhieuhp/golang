package storage

import (
	"context"
	"social-todo-list/modules/item/model"
)

func (sql *sqlStorage) DeleteItem(ctx context.Context, cond map[string]interface{}) error {

	deletedStatus := model.ItemStatusDeleted

	if err := sql.db.Table(model.TodoIem{}.TableName()).
		Where(cond).
		Updates(map[string]interface{}{
			"status": deletedStatus,
		}).Error; err != nil {
		return err
	}

	return nil
}
