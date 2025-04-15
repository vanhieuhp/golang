package service

import (
	"context"
	"errors"
	"log"
	"social-todo-list/common"
	"social-todo-list/modules/userlikeitem/model"
	"social-todo-list/pubsub"
)

type UserLikeItemStore interface {
	Find(ctx context.Context, userId, itemId int) (*model.Like, error)
	Create(ctx context.Context, data *model.Like) error
}

type IncreaseStorage interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeItemService struct {
	store UserLikeItemStore
	//itemStore IncreaseStorage
	ps pubsub.PubSub
}

func NewUserLikeItemService(
	store UserLikeItemStore,
	//itemStore IncreaseStorage,
	ps pubsub.PubSub) *userLikeItemService {
	return &userLikeItemService{
		store: store,
		//itemStore: itemStore,
		ps: ps}
}

func (service *userLikeItemService) LikeItem(ctx context.Context, data *model.Like) error {

	likedItem, err := service.store.Find(ctx, data.UserId, data.ItemId)
	if err != nil && !errors.Is(err, common.RecordNotFound) {
		return model.ErrCannotLikeItem(err)
	}

	if likedItem != nil {
		return nil
	}

	if err := service.store.Create(ctx, data); err != nil {
		return model.ErrCannotLikeItem(err)
	}

	if err := service.ps.Publish(ctx, common.TopicUserLikedItem, pubsub.NewMessage(data)); err != nil {
		log.Println(err)
	}

	//job := asyncjob.NewJob(func(ctx context.Context) error {
	//	if err := service.itemStore.IncreaseLikeCount(ctx, data.ItemId); err != nil {
	//		return err
	//	}
	//
	//	return nil
	//})
	//
	//if err := asyncjob.NewGroup(true, job).Run(ctx); err != nil {
	//	log.Println(err)
	//}

	return nil
}
