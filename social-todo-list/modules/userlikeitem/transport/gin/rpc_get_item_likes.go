package ginuserlikeitem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-todo-list/common"
	"social-todo-list/modules/userlikeitem/storage"
)

func GetItemLikes(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {

		type RequestData struct {
			Ids []int `json:"ids"`
		}

		var data RequestData

		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSqlStorage(db)

		mapRs, err := store.GetItemLikes(context.Request.Context(), data.Ids)

		if err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(mapRs))
	}
}
