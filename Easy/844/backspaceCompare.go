package main

import "fmt"

func backspaceCompare(s string, t string) bool {
	i := len(s) - 1
	j := len(t) - 1
	countS := 0
	countT := 0

	for i >= 0 || j >= 0 {
		for i >= 0 {
			if s[i] == '#' {
				countS++
				i--
			} else if countS > 0 {
				countS--
				i--
			} else {
				break
			}
		}
		for j >= 0 {
			if t[j] == '#' {
				countT++
				j--
			} else if countT > 0 {
				countT--
				j--
			} else {
				break
			}
		}
		if i >= 0 && j >= 0 && s[i] != t[j] {
			return false
		}

		if (i >= 0) != (j >= 0) {
			return false
		}
		i--
		j--
	}
	return true
}

func main() {
	s := "ab#c"
	t := "ad#c"
	fmt.Println(backspaceCompare(s, t))
}