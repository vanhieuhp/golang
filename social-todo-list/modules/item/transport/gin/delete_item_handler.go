package ginitem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-todo-list/common"
	"social-todo-list/modules/item/biz"
	"social-todo-list/modules/item/storage"
	"strconv"
)

func DeleteItem(db *gorm.DB) func(context *gin.Context) {
	return func(context *gin.Context) {

		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		store := storage.NewSqlStorage(db)
		business := biz.NewDeleteItemBiz(store)

		if err := business.DeleteItemById(context.Request.Context(), id); err != nil {
			context.JSON(http.StatusBadRequest, err)

			return
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}