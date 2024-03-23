package main

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	word := strings.Fields(s)
	return len(word[len(word)-1])
}

// func main() {
// 	fmt.Println(lengthOfLastWord("Hello World"))                 // 5
// 	fmt.Println(lengthOfLastWord("   fly me   to   the moon  ")) // 4
// 	fmt.Println(lengthOfLastWord("luffy is still joyboy"))       // 6
// }
