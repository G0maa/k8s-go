package main

import (
	"fmt"
	"os"

	"github.com/g0maa/k8s-go/part1/ex110/random-string-reader/reader"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	router := gin.Default()

	api := router.Group("/api")
	apiV1 := api.Group("/v1")
	
	reader.Register(apiV1.Group("/reader"))

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
