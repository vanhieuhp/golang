package subscriber

import (
	"context"
	"gorm.io/gorm"
	"log"
	"social-todo-list/pubsub"
)

type HasUserId interface {
	GetUserId() int
}

func PushNotificationAfterUserLikeItem(db *gorm.DB) subJob {
	return subJob{
		Title: "push notification after user like item",
		Hld: func(ctx context.Context, message *pubsub.Message) error {

			data := message.Data().(HasUserId)

			log.Println("Noti: ", data.GetUserId())
			return nil
		},
	}
}
