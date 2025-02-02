package main

import "fmt"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}
	count := 0
	for num := range numSet {
		if !numSet[num-1] {
			currNum := num
			currStreak := 1
			for numSet[currNum+1] {
				currNum++
				currStreak++
			}
			if currStreak > count {
				count = currStreak
			}
		}
	}
	return count
}

func main() {
	nums := []int{100,4,200,1,3,2}
	fmt.Println(longestConsecutive(nums))
}