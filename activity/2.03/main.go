package main

import "fmt"

func main() {
	ints := []int{5, 8, 2, 4, 0, 1, 3, 7, 9, 6}
	fmt.Println("Before:", ints)

	for swapped := true; swapped; {
		swapped = false
		for i := 1; i < len(ints); i++ {
			if ints[i-1] > ints[i] {
				ints[i], ints[i-1] = ints[i-1], ints[i]
				swapped = true
			}
		}
	}
	fmt.Println("After:", ints)
}
