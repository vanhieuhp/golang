package upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-todo-list/common"
	"time"
)

func Upload(db * gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		fileHeader, err := ctx.FormFile("file")
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		dst := fmt.Sprintf("static/%d.%s", time.Now().UTC().UnixNano(), fileHeader.Filename)
		if err := ctx.SaveUploadedFile(fileHeader, dst); err != nil {

		}

		img := common.Image{
			Id: 0,
			Url: dst,
			Width: 512,
			Height: 512,
			CloudName: "local",
			Extension: "",
		}
		img.Fulfill("http://localhost:8080")
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
