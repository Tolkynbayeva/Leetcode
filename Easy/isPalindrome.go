package main

import "fmt"

func isPalindrome(s string) bool {
	low := ""
    for _, v := range s {
			if (v >= 'A' && v <= 'Z') || (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {
			 low += isLower(v)
			}
		}
		left := 0
		right := len(low)-1
		for left < right {
			if low[left] == low[right] {
				left++
				right--
			} else {
				return false
			}
		}
		return true
}

func isLower(ch rune) string {
	if ch >= 'A' && ch <= 'Z' {
		ch = ch + 32
	}
	return string(ch)
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println(isPalindrome("race a car")) // false
	fmt.Println(isPalindrome(" ")) // true
}