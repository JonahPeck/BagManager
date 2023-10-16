package main

import (
	"BagManager/controllers"
	"BagManager/models"

	"github.com/gin-gonic/gin"
)

// hello

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.POST("/posts", controllers.CreatePost)
	router.GET("/posts", controllers.FindPosts)

	router.Run("localhost:8080")
}
