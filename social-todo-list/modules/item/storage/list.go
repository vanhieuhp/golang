package storage

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

func (sql *sqlStorage) ListItem(ctx context.Context, filter *model.Filter, paging *common.Paging, moreKeys ...string) ([]model.TodoIem, error) {

	var result []model.TodoIem

	db := sql.db.Where("status <> ?", "Deleted")

	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}

	if err := db.Table(model.TodoIem{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if v := paging.FakeCursor; v != "" {
		uid, err := common.DecodeUID(v)

		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	db = db.Preload("Owner")

	if err := db.Order("id desc").
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}

	if len(result) > 0 {
		result[len(result)-1].Mask()
		paging.NextCursor = result[len(result)-1].FakeId
	}

	return result, nil
}
