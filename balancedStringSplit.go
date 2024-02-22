package main

import "fmt"

func balancedStringSplit(s string) int {
	countL := 0
	balance := 0
	for _, l := range s {
		if l == 76 {
			countL++
		} else {
			countL--
		}
		if countL == 0 {
			balance++
		}
	}
	return balance
}

func main() {
	fmt.Println(balancedStringSplit("RLRRLLRLRL")) //4
	fmt.Println(balancedStringSplit("RLRRRLLRLL")) //2
	fmt.Println(balancedStringSplit("LLLLRRRR"))   //1
}
