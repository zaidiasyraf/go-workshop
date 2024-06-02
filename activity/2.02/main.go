package main

import "fmt"

func main() {
	words := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}

	highestV := 0
	w := ""
	for k, v := range words {
		if v > highestV {
			w = k
			highestV = v
		}
	}
	fmt.Println("Most popular word:", w)
	fmt.Println("With a count of:", highestV)
}
