package ginuser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-todo-list/common"
	"social-todo-list/component/tokenprovider"
	"social-todo-list/modules/user/model"
	"social-todo-list/modules/user/service"
	"social-todo-list/modules/user/storage"
)

func Login(db *gorm.DB, tokenProvider tokenprovider.Provider) func(*gin.Context) {
	return func(c *gin.Context) {
		var loginUserData model.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewSqlStorage(db)
		md5 := common.NewMD5Hash()
		expiry := 60*60*24*30

		service := service.NewLoginService(store, tokenProvider, md5, expiry)
		account, err := service.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))

	}
}