package main

import (
	"strings"
)

func truncateSentence(s string, k int) string {
	result := ""
	array := strings.Fields(s)

	for i, v := range array {
		if i == k {
			break
		}
		result += v + " "
	}
	result = strings.TrimSpace(result)

	return result
}

// func main() {
// 	fmt.Println(truncateSentence("Hello how are you Contestant", 4)) //"Hello how are you"
// 	fmt.Println(truncateSentence("What is the solution to this problem", 4)) //"What is the solution"
// 	fmt.Println(truncateSentence("chopper is not a tanuki", 5)) //"chopper is not a tanuki"
// }
