package handler

import (
	"myapp/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func RegisterUserRoutes(r *gin.Engine, db *gorm.DB) {
	h := &Handler{DB: db}

	r.POST("/login", h.Login)
	r.POST("/register_user", h.RegisterUser)

	//認証必須のグループ
	authorized_r := r.Group("/auth")
	authorized_r.Use(middleware.AuthMiddleware())
	{
		authorized_r.GET("/hello", hello_world)
		authorized_r.POST("/upload_give_item", h.UploadGiveItem)
	}
}

func hello_world(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello world"})
}
