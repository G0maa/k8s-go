package reader

import (
	"bufio"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var random = uuid.New()
var r *bufio.Reader


func init() {
	var err error
	file, err := os.Open("/usr/src/app/output.log")
	if err != nil {
		fmt.Println(err)
	}
	r = bufio.NewReader(file)
}

func reader(c *gin.Context) {
	line, err := r.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"line":  line,
		"random": random,
	})
}