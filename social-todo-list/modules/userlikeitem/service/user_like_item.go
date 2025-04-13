package service

import (
	"context"
	"log"
	"social-todo-list/common"
	"social-todo-list/modules/userlikeitem/model"
)

type UserLikeItemStore interface {
	Create(ctx context.Context, data *model.Like) error
}

type IncreaseStorage interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeItemService struct {
	store     UserLikeItemStore
	itemStore IncreaseStorage
}

func NewUserLikeItemService(store UserLikeItemStore, itemStore IncreaseStorage) *userLikeItemService {
	return &userLikeItemService{store: store, itemStore: itemStore}
}

func (service *userLikeItemService) LikeItem(ctx context.Context, data *model.Like) error {
	if err := service.store.Create(ctx, data); err != nil {
		return model.ErrCannotLikeItem(err)
	}

	go func() {
		defer common.Recovery()
		if err := service.itemStore.IncreaseLikeCount(ctx, data.ItemId); err != nil {
			log.Println(err)
		}
	}()

	return nil
}
