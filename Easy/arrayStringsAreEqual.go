package main

import (
	"strings"
)

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	wordString1 := strings.Join(word1, "")
	wordString2 := strings.Join(word2, "")

	if wordString1 == wordString2 {
		return true
	} else {
		return false
	}
}

// func main() {
// 	fmt.Println(arrayStringsAreEqual([]string{"ab", "c"}, []string{"a", "bc"}))           //true
// 	fmt.Println(arrayStringsAreEqual([]string{"a", "cb"}, []string{"ab", "c"}))           //false
// 	fmt.Println(arrayStringsAreEqual([]string{"abc", "d", "defg"}, []string{"abcddefg"})) //true
// }
