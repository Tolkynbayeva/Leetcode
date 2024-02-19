package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	runes := []rune(s)
	result := ""
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	sentence := strings.Fields(string(runes))
	for _, words := range sentence {
		result = words + " " + result
	}
	result = strings.TrimSpace(result)
	return result
}

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
	fmt.Println(reverseWords("Mr Ding"))
}
