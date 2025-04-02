package service

import (
	"context"
	"social-todo-list/modules/item/model"
)

type DeleteItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoIem, error)
	DeleteItem(ctx context.Context, cond map[string]interface{}) error
}

type deleteItemBiz struct {
	store DeleteItemStorage
}

func NewDeleteItemBiz(store DeleteItemStorage) *deleteItemBiz {
	return &deleteItemBiz{store: store}
}

func (biz *deleteItemBiz) DeleteItemById(ctx context.Context, id int) error {

	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if data.Status != nil && *data.Status == model.ItemStatusDeleted {
		return model.ErrItemIdDeleted
	}

	if err := biz.store.DeleteItem(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}
	return nil
}
