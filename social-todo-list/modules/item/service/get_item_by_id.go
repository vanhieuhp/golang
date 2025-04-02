package service

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

type GetItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoIem, error)
}

type getItemBiz struct {
	store GetItemStorage
}

func NewGetItemBiz(store GetItemStorage) *getItemBiz {
	return &getItemBiz{store: store}
}

func (biz *getItemBiz) GetItemById(ctx context.Context, id int) (*model.TodoIem, error) {

	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(model.EntityName, err)
	}

	return data, nil
}
