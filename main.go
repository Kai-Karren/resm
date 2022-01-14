package main

import (
	"github.com/Kai-Karren/resm/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/request", api.HandleRequest)

	router.Run("localhost:8080")
}
