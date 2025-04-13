package ginitem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-todo-list/common"
	"social-todo-list/modules/item/service"
	"social-todo-list/modules/item/storage"
)

func GetItem(db *gorm.DB) func(context *gin.Context) {
	return func(context *gin.Context) {
		id, err := common.DecodeUID(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		store := storage.NewSqlStorage(db)
		business := service.NewGetItemBiz(store)

		data, err := business.GetItemById(context.Request.Context(), int(id.GetLocalID()))
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
