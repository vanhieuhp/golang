package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

type UpdateItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoIem, error)
	UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TodoItemUpdate) error
}

type updateItemService struct {
	store UpdateItemStorage
	requester common.Requester
}

func NewUpdateItemBiz(store UpdateItemStorage, request common.Requester) *updateItemService {
	return &updateItemService{store: store, requester: request}
}

func (service *updateItemService) UpdateItemById(ctx context.Context, id int, dataUpdate *model.TodoItemUpdate) error {

	data, err := service.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return common.ErrCannotGetEntity(model.EntityName, err)
		}
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	isOwner := service.requester.GetUserId() == data.UserId

	if !isOwner && !common.IsAdmin(service.requester) {
		return common.ErrNoPermission(errors.New("No permission"))
	}

	if data.Status != nil && *data.Status == model.ItemStatusDeleted {
		return common.ErrEntityDeleted(model.EntityName, model.ErrItemIdDeleted)
	}

	if err := service.store.UpdateItem(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return err
	}
	return nil
}
