package main

import "fmt"

func numJewelsInStones(jewels string, stones string) int {
	count := 0
	for _, chJewels := range jewels {
		for _, chStones := range stones {
			if chJewels == chStones {
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
	fmt.Println(numJewelsInStones("z", "ZZ"))
}
