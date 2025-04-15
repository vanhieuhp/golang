package subscriber

import (
	"context"
	"gorm.io/gorm"
	"social-todo-list/modules/item/storage"
	"social-todo-list/pubsub"
)

type HasItemId interface {
	GetItemId() int
}

//func IncreaseLikeCountAfterUserLikeItem(db *gorm.DB, serviceCtx goservice.ServiceContext, ctx context.Context) {
//	ps := serviceCtx.MustGet(common.PluginPubSub).(pubsub.PubSub)
//	c, _ := ps.Subscribe(ctx, common.TopicUserLikedItem)
//
//	go func() {
//		defer common.Recovery()
//
//		for msg := range c { // msg := <-c
//
//			data := msg.Data().(HasItemId)
//
//			if err := storage.NewSqlStorage(db).IncreaseLikeCount(ctx, data.GetItemId()); err != nil {
//				log.Println(err)
//			}
//			log.Println("Increase Like Count After User Liked Item: ", msg.Data)
//		}
//
//	}()
//}

func IncreaseLikeCountAfterUserLikeItem(db *gorm.DB) subJob {
	return subJob{
		Title: "Increase like count after user like item",
		Hld: func(ctx context.Context, message *pubsub.Message) error {

			data := message.Data().(HasItemId)
			return storage.NewSqlStorage(db).IncreaseLikeCount(ctx, data.GetItemId())
		},
	}
}
