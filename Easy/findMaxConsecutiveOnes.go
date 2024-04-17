package main

import (
	"fmt"
)

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			if count > max {
				max = count
			}
			count = 0

		}
	}
	if count > max {
		max = count
	}

	return max
}

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1})) // 3

	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1})) // 3
	//	fmt.Println(findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1})) // 2

}
