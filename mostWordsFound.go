package main

import (
	"fmt"
	"strings"
)

func mostWordsFound(sentences []string) int {
	maxCountWords := 0
	for _, sentence := range sentences {
		newSentence := strings.Fields(sentence)
		countWords := len(newSentence)
		if countWords > maxCountWords {
			maxCountWords = countWords
		}
	}
	return maxCountWords
}

func main() {
	case1 := mostWordsFound([]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}) //6
	case2 := mostWordsFound([]string{"please wait", "continue to fight", "continue to win"})  //3
	fmt.Println(case1)
	fmt.Println(case2)

}
