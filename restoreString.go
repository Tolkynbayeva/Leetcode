package main

import (
	"fmt"
)

func restoreString(s string, indices []int) string {
    result := make([]rune, len(s))

	for i := 0; i < len(s); i++ {
		if indices[i] >= 0 && indices[i] < len(s) {
			result[indices[i]] = rune(s[i])
		}
	}
	return string(result)
}

func main() {
	fmt.Println(restoreString("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3})) //leetcode
	fmt.Println(restoreString("abc", []int{0, 1, 2})) //abc
}
