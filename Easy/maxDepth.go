package main

import "fmt"

func maxDepth(s string) int {
	stack := []rune{}
	result := 0
	for _, v := range s {
		if v == '(' {
			stack = append(stack, v)
			if len(stack) > result {
				result = len(stack)
			}
		} else if v == ')' {
			stack = stack[:len(stack)-1]
		}
	}
	return result
}

func main() {
	//fmt.Println(maxDepth("(1+(2*3)+((8)/4))+1")) // 3
	//fmt.Println(maxDepth("(1)+((2))+(((3)))"))   // 3
	fmt.Println(maxDepth("8*((1*(5+6))*(8/6))")) // 3
}
