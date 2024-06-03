package main

import (
	"fmt"
	"unicode"
)

func main() {
	if passwordhecker("") {
		fmt.Println("Password good")
	} else {
		fmt.Println("Password bad")
	}

	if passwordhecker("This!I5A") {
		fmt.Println("Password good")
	} else {
		fmt.Println("Password bad")
	}
}

func passwordhecker(pw string) bool {
	pwR := []rune(pw)

	if len(pwR) < 8 {
		return false
	}
	if len(pwR) > 15 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	for _, v := range pwR {
		if unicode.IsUpper(v) {
			hasUpper = true
		}
		if unicode.IsLower(v) {
			hasLower = true
		}
		if unicode.IsNumber(v) {
			hasNumber = true
		}
		if unicode.IsPunct(v) || unicode.IsSymbol(v) {
			hasSymbol = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSymbol
}