package ginuser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-todo-list/common"
	"social-todo-list/modules/user/model"
	"social-todo-list/modules/user/service"
	"social-todo-list/modules/user/storage"
)

func Register(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		var data model.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		store := storage.NewSqlStorage(db)
		md5 := common.NewMD5Hash()
		service := service.NewRegisterService(store, md5)

		if err := service.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))

	}
}

