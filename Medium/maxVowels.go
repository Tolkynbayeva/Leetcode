package main

import "fmt"

func maxVowels(s string, n int) int {
	count := 0
	max := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "a" || string(s[i]) == "e" || string(s[i]) == "o" || string(s[i]) == "u" || string(s[i]) == "i" {
			count++
		}
		if i >= n {
			if string(s[i-n]) == "a" || string(s[i-n]) == "e" || string(s[i-n]) == "o" || string(s[i-n]) == "u" || string(s[i-n]) == "i" {
				count--
			}
		}
		if count > max {
			max = count
		}
	}
	return max
}

func main() {
	fmt.Println(maxVowels("abciiidef", 3)) // 3
	// fmt.Println(maxVowels("aeiou", 2))     // 2
	// fmt.Println(maxVowels("leetcode", 3))  // 2
}
