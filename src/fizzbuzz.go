package cmlabs

import "fmt"

func Fizz(number int) bool {
	return number%3 == 0
}

func Buzz(number int) bool {
	return number%5 == 0
}

func Fizzbuzz(n int) {
	for num := 1; num <= n; num++ {
		newLine := false

		if Fizz(num) {
			fmt.Print("Fizz")
			newLine = true
		}

		if Buzz(num) {
			fmt.Print("Buzz")
			newLine = true
		}

		if newLine {
			fmt.Println()
		}
	}
}
