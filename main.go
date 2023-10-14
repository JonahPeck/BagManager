package main

import (
	"JonahPeck/BagManager/BagManager/models"

	"github.com/gin-gonic/gin"
)

// hello

func main() {
	router := gin.Default()
	models.ConnectDatabase()

	router.Run("localhost:8080")
}
