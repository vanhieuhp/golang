package cmd

import (
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"social-todo-list/common"
	"social-todo-list/component/jwt"
	"social-todo-list/middleware"
	ginitem "social-todo-list/modules/item/transport/gin"
	authStore "social-todo-list/modules/user/storage"
	ginuser "social-todo-list/modules/user/transport/gin"
	ginuserlikeitem "social-todo-list/modules/userlikeitem/transport/gin"
	"social-todo-list/pubsub"
	"social-todo-list/subscriber"
	"social-todo-list/upload"
)

func newService() goservice.Service {
	service := goservice.New(
		goservice.WithName("social-todo-list"),
		goservice.WithVersion("1.0.0"),
		goservice.WithInitRunnable(pubsub.NewPubSub(common.PluginPubSub)),
		//goservice.WithInitRunnable(simple.NewSimplePlugin("simple")),
	)

	return service
}

func setupDatabase() (*gorm.DB, error) {
	dsn := os.Getenv("DB_CONN")
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func setupAuth(db *gorm.DB) (*jwt.JwtProvider, func(c *gin.Context)) {
	systemSecret := os.Getenv("JWT_SECRET")
	tokenProvider := jwt.NewTokenJWTProvider("jwt", systemSecret)
	authStorage := authStore.NewSqlStorage(db)
	return tokenProvider, middleware.RequiredAuth(authStorage, tokenProvider)
}

func registerRoutes(service goservice.Service, db *gorm.DB, tokenProvider *jwt.JwtProvider, middleAuth func(c *gin.Context)) {
	service.HTTPServer().AddHandler(func(engine *gin.Engine) {
		engine.Use(middleware.Recovery())

		//if val, ok := service.MustGet("simple").(interface{ GetValue() string }); ok {
		//	log.Println(val.GetValue())
		//} else {
		//	log.Println("simple plugin missing or invalid")
		//}

		v1 := engine.Group("/api/v1")
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

				items.POST("/:id/like", middleAuth, ginuserlikeitem.LikeItem(service, db))
				items.DELETE("/:id/unlike", middleAuth, ginuserlikeitem.UnlikeItem(service, db))
				items.GET("/:id/liked-users", middleAuth, ginuserlikeitem.ListUserLikedItem(db))
			}
		}

		engine.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})

		engine.Run("localhost:8080")
	})
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start social Todo Service",
	Run: func(cmd *cobra.Command, args []string) {
		service := newService()

		logger := service.Logger("service")

		if err := service.Init(); err != nil {
			logger.Fatal("Service init failed: ", err)
		}

		db, err := setupDatabase()
		if err != nil {
			logger.Fatal("DB setup failed: ", err)
		}

		tokenProvider, middleAuth := setupAuth(db)
		registerRoutes(service, db, tokenProvider, middleAuth)

		subscriber.NewEngine(service, db).Start()

		if err := service.Start(); err != nil {
			logger.Fatal("Service failed to start: ", err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Command execution failed: %v", err)
	}
}
