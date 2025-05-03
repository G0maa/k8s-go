package reader

import (
	"github.com/gin-gonic/gin"
)

func Register(router *gin.RouterGroup) {
	router.GET("/", reader)
}