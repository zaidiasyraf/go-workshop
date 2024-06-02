package main

import "fmt"

// fizz buzz
func main() {

	for i := 0; i < 100; i++ {
		j := i + 1
		if j%3 == 0 && j%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if j%3 == 0 {
			fmt.Println("Fizz")
		} else if j%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(j)
		}
	}

}
