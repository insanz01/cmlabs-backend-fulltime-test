package cmlabs_test

import (
	"testing"

	cmlabs "github.com/insanz01/cmlabs-backend-fulltime-test/src"
)

func TestFizzBuzz(t *testing.T) {
	fizzBuzzTest := []struct {
		Number int
		Method string
		Expect bool
	}{
		{
			Number: 3,
			Method: "Fizz",
			Expect: true,
		},
		{
			Number: 5,
			Method: "Buzz",
			Expect: true,
		},
		{
			Number: 15,
			Method: "FizzBuzz",
			Expect: true,
		},
		{
			Number: 20,
			Method: "Fizz",
			Expect: false,
		},
		{
			Number: 24,
			Method: "Buzz",
			Expect: false,
		},
		{
			Number: 30,
			Method: "FizzBuzz",
			Expect: true,
		},
	}

	for _, fizzBuzz := range fizzBuzzTest {
		switch fizzBuzz.Method {
		default:
			t.Errorf("Invalid method (%s)", fizzBuzz.Method)
		case "Fizz":
			result := cmlabs.Fizz(fizzBuzz.Number)
			if result != fizzBuzz.Expect {
				t.Errorf("Hasil (%t) tidak sesuai dengan yang diharapkan (%t)", result, fizzBuzz.Expect)
			}
		case "Buzz":
			result := cmlabs.Buzz(fizzBuzz.Number)
			if result != fizzBuzz.Expect {
				t.Errorf("Hasil (%t) tidak sesuai dengan yang diharapkan (%t)", result, fizzBuzz.Expect)
			}
		case "FizzBuzz":
			fizzResult := cmlabs.Fizz(fizzBuzz.Number)
			buzzResult := cmlabs.Buzz(fizzBuzz.Number)

			if fizzResult == buzzResult {
				if fizzResult != fizzBuzz.Expect {
					t.Errorf("Hasil (%t) tidak sesuai dengan yang diharapkan (%t)", fizzResult, fizzBuzz.Expect)
				}
			}
		}
	}
}
