package main

import (
	"fmt"
	"os"
	"time"
)


func main() {
	f, err := os.Create("/usr/src/app/output.log")
	
	if err != nil {
		fmt.Println(err)
	}

	// After a bit of googling it seems this won't run with Ctrl+C.
	// Will need t do some go magic: https://stackoverflow.com/questions/50406639/are-deferred-functions-called-when-sigint-is-received-in-go
	defer f.Close()

	for {
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("Writing timestamp to file... %s\n", timestamp)

		f.WriteString(fmt.Sprintf("%s\n", timestamp))
		
		time.Sleep(time.Second * 5)
	}

}