package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"social-todo-list/common"
	"social-todo-list/component/jwt"
	"social-todo-list/middleware"
	"social-todo-list/modules/item/transport/gin"
	"social-todo-list/modules/user/storage"
	ginuser "social-todo-list/modules/user/transport/gin"
	ginuserlikeitem "social-todo-list/modules/userlikeitem/transport/gin"
	"social-todo-list/upload"
)

func main() {
	dsn := os.Getenv("DB_CONN")
	systemSecret := os.Getenv("JWT_SECRET")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	tokenProvider := jwt.NewTokenJWTProvider("jwt", systemSecret)
	authStore := storage.NewSqlStorage(db)
	middleAuth := middleware.RequiredAuth(authStore, tokenProvider)

	router := gin.Default()
	router.Use(middleware.Recovery())

	v1 := router.Group("/api/v1")
	{
		v1.PUT("/upload", upload.Upload(db))

		v1.POST("/register", ginuser.Register(db))
		v1.POST("/login", ginuser.Login(db, tokenProvider))
		v1.GET("/profile", middleAuth, ginuser.Profile())

		items := v1.Group("/items")
		{
			items.POST("", middleAuth, ginitem.CreateItem(db))
			items.GET("", middleAuth, ginitem.ListItem(db))
			items.GET("/:id", ginitem.GetItem(db))
			items.PUT("/:id", middleAuth, ginitem.UpdateItem(db))
			items.DELETE("/:id", middleAuth, ginitem.DeleteItem(db))

			items.POST("/:id/like", middleAuth, ginuserlikeitem.LikeItem(db))
			items.DELETE("/:id/unlike", middleAuth, ginuserlikeitem.UnlikeItem(db))
			items.GET("/:id/liked-users", middleAuth, ginuserlikeitem.ListUserLikedItem(db))
		}
	}

	router.GET("/ping", func(c *gin.Context) {

		go func() {
			defer common.Recovery()
			fmt.Println([]int{}[0])
		}()
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err = router.Run("localhost:8080")
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
