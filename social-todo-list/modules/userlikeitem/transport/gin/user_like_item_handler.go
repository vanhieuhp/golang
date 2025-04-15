package ginuserlikeitem

import (
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-todo-list/common"
	"social-todo-list/modules/userlikeitem/model"
	"social-todo-list/modules/userlikeitem/service"
	"social-todo-list/modules/userlikeitem/storage"
	"social-todo-list/pubsub"
)

func LikeItem(serviceCtx goservice.ServiceContext, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := common.DecodeUID(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		//db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)
		ps := serviceCtx.MustGet(common.PluginPubSub).(pubsub.PubSub)

		store := storage.NewSqlStorage(db)
		//itemStore := itemStore.NewSqlStorage(db)
		userLikeItemService := service.NewUserLikeItemService(store, ps)

		if err := userLikeItemService.LikeItem(c.Request.Context(), &model.Like{
			UserId:    requester.GetUserId(),
			ItemId:    int(id.GetLocalID()),
			CreatedAt: nil,
		}); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
