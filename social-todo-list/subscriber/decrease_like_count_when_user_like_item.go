package subscriber

import (
	"context"
	"gorm.io/gorm"
	"social-todo-list/modules/item/storage"
	"social-todo-list/pubsub"
)

func DecreaseLikeCountAfterUserLikeItem(db *gorm.DB) subJob {
	return subJob{
		Title: "Increase like count after user like item",
		Hld: func(ctx context.Context, message *pubsub.Message) error {

			data := message.Data().(HasItemId)
			return storage.NewSqlStorage(db).DecreaseLikeCount(ctx, data.GetItemId())
		},
	}
}
