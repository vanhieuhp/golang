package service

import (
	"context"
	"errors"
	"log"
	"social-todo-list/common"
	"social-todo-list/modules/userlikeitem/model"
	"social-todo-list/pubsub"
)

type UserUnlikeItemStore interface {
	Find(ctx context.Context, userId, itemId int) (*model.Like, error)
	Delete(ctx context.Context, userid, itemId int) error
}

type userUnlikeItemService struct {
	store UserUnlikeItemStore
	//itemStore DecreaseItemStorage
	ps pubsub.PubSub
}

//type DecreaseItemStorage interface {
//	DecreaseLikeCount(ctx context.Context, id int) error
//}

func NewUserUnlikeItemService(
	store UserUnlikeItemStore,
	//itemStore DecreaseItemStorage,
	pubSub pubsub.PubSub,
) *userUnlikeItemService {
	return &userUnlikeItemService{
		store: store,
		//itemStore: itemStore,
		ps: pubSub}
}

func (service *userUnlikeItemService) UnlikeItem(ctx context.Context, userid, itemId int) error {
	_, err := service.store.Find(ctx, userid, itemId)

	// Delete if data existed
	if errors.Is(err, common.RecordNotFound) {
		return model.ErrDidNotLikeItem(err)
	}

	if err != nil {
		return model.ErrCannotUnlikeItem(err)
	}

	if err := service.store.Delete(ctx, userid, itemId); err != nil {
		return model.ErrCannotLikeItem(err)
	}

	if err := service.ps.Publish(ctx, common.TopicUserUnLikedItem, pubsub.NewMessage(&model.Like{UserId: userid, ItemId: itemId})); err != nil {
		log.Println(err)
	}

	return nil
}
