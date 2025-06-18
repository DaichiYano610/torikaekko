package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", hello_world)

	router.Run(":8080")
}

func hello_world(c *gin.Context) {

}
