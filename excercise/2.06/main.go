package main

import (
	"fmt"
	"time"
)

func main() {
	dayBorn := time.Sunday

	switch dayBorn {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Born on weekday")
	case time.Saturday, time.Sunday:
		fmt.Println("Born on weekend")
	default:
		fmt.Println("Weird")
	}
}
