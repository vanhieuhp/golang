package ginuserlikeitem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-todo-list/common"
	itemStore "social-todo-list/modules/item/storage"
	"social-todo-list/modules/userlikeitem/model"
	"social-todo-list/modules/userlikeitem/service"
	"social-todo-list/modules/userlikeitem/storage"
)

func LikeItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := common.DecodeUID(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := storage.NewSqlStorage(db)
		itemStore := itemStore.NewSqlStorage(db)
		userLikeItemService := service.NewUserLikeItemService(store, itemStore)

		if err := userLikeItemService.LikeItem(c.Request.Context(), &model.Like{
			UserId:    requester.GetUserId(),
			ItemId:    int(id.GetLocalID()),
			CreatedAt: nil,
		}); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
