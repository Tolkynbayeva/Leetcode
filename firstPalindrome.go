package main

import (
	"fmt"
)

func firstPalindrome(words []string) string {
	palindromic := ""
	reverse := ""

	for _, v := range words {
		if len(palindromic) > 0 {
			break
		}

		reverse = ""
		reverse += v

		runes := []rune(reverse)

		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}

		if reverse == string(runes) {
			palindromic = reverse
		}
	}
	return palindromic
}

func main() {
	fmt.Println(firstPalindrome([]string{"abc", "car", "ada", "racecar", "cool"})) // "ada"
	fmt.Println(firstPalindrome([]string{"notapalindrome", "racecar"}))            // "racecar"
	fmt.Println(firstPalindrome([]string{"def", "ghi"}))                           // ""
}
