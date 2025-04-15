package ginuserlikeitem

import (
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-todo-list/common"
	"social-todo-list/modules/userlikeitem/service"
	"social-todo-list/modules/userlikeitem/storage"
	"social-todo-list/pubsub"
)

func UnlikeItem(serviceCtx goservice.ServiceContext, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := common.DecodeUID(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := storage.NewSqlStorage(db)
		//itemStore := itemStore.NewSqlStorage(db)
		ps := serviceCtx.MustGet(common.PluginPubSub).(pubsub.PubSub)

		userLikeItemService := service.NewUserUnlikeItemService(store, ps)

		if err := userLikeItemService.UnlikeItem(
			c.Request.Context(),
			requester.GetUserId(),
			int(id.GetLocalID()),
		); err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
