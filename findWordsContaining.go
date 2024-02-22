package main

import (
	"fmt"
)

func findWordsContaining(words []string, x byte) []int {
	result := []int{}
	runes := rune(x)

	for i, word := range words {
		for _, char := range word {
			if char == runes {
				result = append(result, i)
				break
			}
		}
	}

	return result
}

func main() {
	fmt.Println(findWordsContaining([]string{"leet", "code"}, byte('e')))  // [0,1]
	fmt.Println(findWordsContaining([]string{"abc", "bcd", "aaaa", "cbc"}, byte('a'))) // [0,2]
	fmt.Println(findWordsContaining([]string{"abc", "bcd", "aaaa", "cbc"}, byte('z'))) // []
}
