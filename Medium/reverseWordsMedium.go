package main

import (
	"strings"
)

func reverseWords(s string) string {
	runes := []rune(s)
	result := ""

	sentence := strings.Fields(string(runes))
	for _, words := range sentence {
		result = words + " " + result
	}
	result = strings.TrimSpace(result)
	return result
}

// func main() {
// 	fmt.Println(reverseWords("the sky is blue"))  // "blue is sky the"
// 	fmt.Println(reverseWords("  hello world  "))  // "world hello"
// 	fmt.Println(reverseWords("a good   example")) // "example good a"
// }
