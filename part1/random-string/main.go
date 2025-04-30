package main

import (
	"fmt"
	"os"

	"github.com/g0maa/k8s-go/part1/random-string/health"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	router := gin.Default()

	health.Register(router.Group("/health"))

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}