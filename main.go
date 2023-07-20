package main

import (
	"fmt"

	cmlabs "github.com/insanz01/cmlabs-backend-fulltime-test/src"
)

func main() {
	fmt.Println("Anda dapat menjalankan dengan go test ./test")

	cmlabs.Fizzbuzz(100)

	word := "KASUR RUSAK"
	palindromeResult := fmt.Sprintf("Check Palindrome (%s) => (%t)", word, cmlabs.Palindrome(word))

	fmt.Println(palindromeResult)

	fmt.Println("Anda dapat menjalankan dengan go test ./test")
}
