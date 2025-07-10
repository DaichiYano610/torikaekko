package main

import (
	"myapp/internal/handler"
	"myapp/internal/initializer"

	"github.com/gin-gonic/gin"
)

func main() {
	initializer.Env_load()
	database := initializer.InitDB()
	router := gin.Default()
	handler.RegisterUserRoutes(router, database)
	router.Run(":8080")
}
