package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)


func main() {
	uuid := uuid.New()
	for {
		fmt.Println("UUID:", uuid)

		time.Sleep(time.Second)
	}
}