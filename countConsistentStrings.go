package main

import "fmt"

func countConsistentStrings(allowed string, words []string) int {
	result := 0

	allowedMap := make(map[rune]bool)
	for _, char := range allowed {
		allowedMap[char] = true
	}
	//fmt.Println(allowedMap)

outerLoop:
	for _, word := range words {
		for _, char := range word {
			if !allowedMap[char] {
				continue outerLoop
			}
		}
		result++
	}

	return result
}

func main() {
	fmt.Println(countConsistentStrings("ab", []string{"ad", "bd", "aaab", "baa", "badab"})) // 2
	fmt.Println(countConsistentStrings("abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}))          // 7
	fmt.Println(countConsistentStrings("cad", []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"})) // 4
}
