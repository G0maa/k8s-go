package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func init() {
	_ = godotenv.Load()
}

func main() {
	counter := 0

	router := gin.Default()
	router.GET("/pingpong", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("pong %d", counter),
		})
		counter++
	})

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}