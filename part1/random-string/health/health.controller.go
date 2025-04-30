package health

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var random = uuid.New()
var timestamp = time.Now().Format("2006-01-02 15:04:05")

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"random":  random,
		"time":    timestamp,
	})
}