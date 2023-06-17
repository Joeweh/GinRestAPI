package main

import (
	"os"
	"log"
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

	port := os.Getenv("PORT")
	
	if port == "" {
		port = "8080"
	}
	
	if err := router.Run(":" + port); err != nil {
        log.Panicf("error: %s", err)
	}
}
