package main

import "fmt"

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	result := 0
	for _, row := range items {
		//fmt.Println("row", row)
			if ruleKey == "type" && ruleValue == row[0] || ruleKey == "color" && ruleValue == row[1] || ruleKey == "name" && ruleValue == row[2]{
				result++
			}
	}
	return result
}

func main() {
	case1 := countMatches([][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}}, "color", "silver") //1
	case2 := countMatches([][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "phone"}, {"phone", "gold", "iphone"}}, "type", "phone") //2
	fmt.Println(case1)
	fmt.Println(case2)
}
