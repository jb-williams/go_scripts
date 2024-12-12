package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current time
	// 	currentTime := time.Now()

	// Format the current time in a human-readable way
	fmt.Println(time.Now().Format("Mon, 01/02/2006, 3:04 PM"))
}

