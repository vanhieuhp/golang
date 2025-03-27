package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"social-todo-list/common"
	"social-todo-list/middleware"
	"social-todo-list/modules/item/transport/gin"
)

func main() {
	dsn := "admin:password@tcp(127.0.0.1:3306)/springboot?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	router := gin.Default()

	router.Use(middleware.Recovery())

	v1 := router.Group("/api/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", ginitem.CreateItem(db))
			items.GET("", ginitem.ListItem(db))
			items.GET("/:id", ginitem.GetItem(db))
			items.PUT("/:id", ginitem.UpdateItem(db))
			items.DELETE("/:id", ginitem.DeleteItem(db))
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

