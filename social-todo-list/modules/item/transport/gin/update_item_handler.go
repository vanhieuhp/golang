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

func UpdateItem(db *gorm.DB) func(context *gin.Context) {
	return func(context *gin.Context) {
		var data model.TodoItemUpdate

		id, err := common.DecodeUID(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		requester := context.MustGet(common.CurrentUser).(common.Requester)
		store := storage.NewSqlStorage(db)
		business := service.NewUpdateItemBiz(store, requester)

		if err := business.UpdateItemById(context, int(id.GetLocalID()), &data); err != nil {
			context.JSON(http.StatusInternalServerError, err)

			return
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
