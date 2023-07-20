package cmlabs

import "fmt"

func Fizz(number int) bool {
	return number%3 == 0
}

func Buzz(number int) bool {
	return number%5 == 0
}

func FizzBuzz(number int) bool {
	return number%3 == 0 && number%5 == 0
}

func PrintFizzbuzz(n int) {
	for num := 1; num <= n; num++ {
		if FizzBuzz(num) {
			fmt.Println("FizzBuzz")
		} else if Fizz(num) {
			fmt.Println("Fizz")
		} else if Buzz(num) {
			fmt.Println("Buzz")
		}
	}
}
