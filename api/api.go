package api

import (
	_ "app/api/docs"
	"app/api/handler"
	"app/config"
	"app/pkg/logger"
	"app/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewApi(r *gin.Engine, cfg *config.Config, storage storage.StorageI, logger logger.LoggerI) {

	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization

	handler := handler.NewHandler(cfg, storage, logger)

	r.Use(customCORSMiddleware())

	// Login Api
	r.POST("/login", handler.Login)

	// Register Api
	r.POST("/register", handler.Register)

	// User Api
	r.POST("/user", handler.AuthMiddleware(), handler.CreateUser)
	r.GET("/user/:id", handler.AuthMiddleware(), handler.GetByIdUser)
	r.GET("/user", handler.GetListUser)
	r.PUT("/user/:id", handler.UpdateUser)
	r.DELETE("/user/:id", handler.AuthMiddleware(), handler.DeleteUser)

	r.POST("/question", handler.CreateQuestion)
	r.GET("/question/:id", handler.GetByIdQuestion)
	r.GET("/question", handler.GetListQuestion)
	r.PUT("/question/:id", handler.UpdateQuestion)
	r.DELETE("/question/:id", handler.DeleteQuestion)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Accesp-Encoding, Authorization, Cache-Control")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
