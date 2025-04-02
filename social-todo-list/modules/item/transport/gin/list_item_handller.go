package ginitem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
	"social-todo-list/modules/item/service"
	"social-todo-list/modules/item/storage"
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

		store := storage.NewSqlStorage(db)
		business := service.NewListItemBiz(store)

		result, err := business.ListItem(context.Request.Context(), &filter, &paging)

		if err != nil {
			context.JSON(http.StatusBadRequest, err)
		}

		context.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
