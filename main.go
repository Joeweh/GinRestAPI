package main

import (
	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/users", GetUsers)

	router.POST("/users", SaveUser)

	router.Run("localhost:8080")
}
