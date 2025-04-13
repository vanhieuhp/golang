package ginuserlikeitem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-todo-list/common"
	"social-todo-list/modules/userlikeitem/service"
	"social-todo-list/modules/userlikeitem/storage"
)

func ListUserLikedItem(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var paging common.Paging

		id, err := common.DecodeUID(context.Param("id"))

		if err := context.ShouldBind(&paging); err != nil {
			context.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		paging.Process()
		store := storage.NewSqlStorage(db)
		business := service.NewListUserLikeItemService(store)

		result, err := business.ListUserLikedItem(context.Request.Context(), int(id.GetLocalID()), &paging)

		if err != nil {
			context.JSON(http.StatusBadRequest, err)
		}

		for i := range result {
			result[i].Mask()
		}

		context.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
