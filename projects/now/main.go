package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	fmt.Println("Current Time:", formattedTime)

}
