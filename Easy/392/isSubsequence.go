package main

import "fmt"

func isSubsequence(s string, t string) bool {
	l := 0
	r := 0
	for r < len(t) {
		if l < len(s) && s[l] == t[r] {
			l++
		}
		r++
	}
	return l == len(s)
}

func main() {
	s := "abc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}
