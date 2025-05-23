package ginitem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-todo-list/common"
	"social-todo-list/modules/item/biz"
	"social-todo-list/modules/item/model"
	"social-todo-list/modules/item/storage"
	"strconv"
)

func UpdateItem(db *gorm.DB) func(context *gin.Context) {
	return func(context *gin.Context) {
		var data model.TodoItemUpdate

		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSqlStorage(db)
		business := biz.NewUpdateItemBiz(store)

		if err := business.UpdateItemById(context, id, &data); err != nil {
			context.JSON(http.StatusInternalServerError, err)

			return
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}