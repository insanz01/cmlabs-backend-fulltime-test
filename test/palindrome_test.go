package cmlabs_test

import (
	"testing"

	cmlabs "github.com/insanz01/cmlabs-backend-fulltime-test/src"
)

func TestPalindrome(t *testing.T) {
	palindromeTest := []struct {
		Word   string
		Expect bool
	}{
		{
			Word:   "SAIPPUAKIVIKAUPPIAS",
			Expect: true,
		},
		{
			Word:   "Aibohphobia",
			Expect: true,
		},
		{
			Word:   "Anna",
			Expect: true,
		},
		{
			Word:   "Civic",
			Expect: true,
		},
		{
			Word:   "My gym",
			Expect: true,
		},
		{
			Word:   "No lemon, no melon",
			Expect: true,
		},
		{
			Word:   "go green",
			Expect: false,
		},
	}

	for _, palindrome := range palindromeTest {
		result := cmlabs.Palindrome(palindrome.Word)

		if result != palindrome.Expect {
			t.Errorf("Hasil (%t) tidak sesuai dengan yang diharapkan (%t)", result, palindrome.Expect)
		}
	}
}
