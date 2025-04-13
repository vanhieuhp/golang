package ginitem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
	"social-todo-list/modules/item/repository"
	"social-todo-list/modules/item/service"
	"social-todo-list/modules/item/storage"
	userLikeStore "social-todo-list/modules/userlikeitem/storage"
)

func ListItem(db *gorm.DB) func(context *gin.Context) {
	return func(context *gin.Context) {
		var paging common.Paging

		if err := context.ShouldBind(&paging); err != nil {
			context.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		paging.Process()

		var filter model.Filter
		if err := context.ShouldBind(&filter); err != nil {
			context.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		requester := context.MustGet(common.CurrentUser).(common.Requester)

		store := storage.NewSqlStorage(db)
		likeStore := userLikeStore.NewSqlStorage(db)
		repo := repository.NewListItemReo(store, likeStore, requester)
		business := service.NewListItemService(repo, requester)

		result, err := business.ListItem(context.Request.Context(), &filter, &paging)

		if err != nil {
			context.JSON(http.StatusBadRequest, err)
		}

		for i := range result {
			result[i].Mask()
		}

		context.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
