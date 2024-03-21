package main

import (
	"fmt"
	"regexp"
	"strings"
)

func sortSentence(s string) string {
	result := ""
	sentence := strings.Fields(s)
	m := map[string]string{}
	ints := "123456789"

	for _, words := range sentence {
		for _, char := range words {
			for _, digit := range ints {
				if char == digit {
					m[string(digit)] = words
					break
				}
			}
		}
	}

	for i, newValue := range ints {
		if m[string(newValue)] == "" {
			break
		}
		result += m[string(newValue)]

		if i != len(result)-1 {
			result += " "
		}
	}

	re := regexp.MustCompile("[0-9]")
	final := re.ReplaceAllString(result, "")
	final = strings.TrimSpace(final)

	fmt.Println(m)
	fmt.Println("r ", result)
	fmt.Println("f ", final)
	return final
}

// func main() {
// 	fmt.Print(sortSentence("is2 sentence4 This1 a3"))
// }
