package main

import (
	"myapp/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	handler.RegisterUserRoutes(router)
	router.Run(":8080")
}
