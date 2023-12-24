package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/custom", CustomBind)
	router.Run("localhost:8080")
}
