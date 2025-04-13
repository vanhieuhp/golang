package service

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/userlikeitem/model"
)

type ListUserLikeItemStore interface {
	ListUsers(
		ctx context.Context,
		itemId int,
		paging *common.Paging,
	) ([]common.SimpleUser, error)
}

type listUserLikeItemService struct {
	store ListUserLikeItemStore
}

func NewListUserLikeItemService(store ListUserLikeItemStore) *listUserLikeItemService {
	return &listUserLikeItemService{store: store}
}

func (service *listUserLikeItemService) ListUserLikedItem(
	ctx context.Context,
	itemId int,
	paging *common.Paging,
) ([]common.SimpleUser, error) {

	result, err := service.store.ListUsers(ctx, itemId, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	return result, nil
}
