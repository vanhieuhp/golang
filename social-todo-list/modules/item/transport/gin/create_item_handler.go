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

func CreateItem(db *gorm.DB) func(context *gin.Context) {
	return func(context *gin.Context) {
		var data model.TodoItemCreation
		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		requester := context.MustGet(common.CurrentUser).(common.Requester)
		data.UserId = requester.GetUserId()

		store := storage.NewSqlStorage(db)
		business := service.NewCreateItemBiz(store)

		if err := business.CreateNewItem(context.Request.Context(), &data); err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}