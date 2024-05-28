package main

import (
	"fmt"
)

func main() {
	count := 5
	message := "Not greater than 5"
	if count > 5 {
		message = "Greater than 5"
	}
	fmt.Println(message)
}
