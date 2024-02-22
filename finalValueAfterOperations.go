package main

import "fmt"

func finalValueAfterOperations(operations []string) int {
    result := 0
	
	for _, x := range operations {
		if x == "X++" || x == "++X" {
			result++
		} else {
			result--
		}
	}

	return result
}

func main() {
	fmt.Println(finalValueAfterOperations([]string{"--X","X++","X++"})) //1
	fmt.Println(finalValueAfterOperations([]string{"++X","++X","X++"})) //3
	fmt.Println(finalValueAfterOperations([]string{"X++","++X","--X","X--"})) //0
}