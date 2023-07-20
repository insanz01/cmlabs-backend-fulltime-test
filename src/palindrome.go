package cmlabs

import "strings"

func Palindrome(text string) bool {
	// Menghilangkan spasi dan mengubah semua karakter menjadi huruf kecil
	text = strings.ReplaceAll(text, " ", "")
	text = strings.ToLower(text)

	// Menguji apakah string adalah palindrome
	for i := 0; i < len(text)/2; i++ {
		if text[i] != text[len(text)-1-i] {
			return false
		}
	}
	return true
}
