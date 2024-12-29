package main

import "fmt"

func sortedSquares(nums []int) []int {
	l := 0
	r := len(nums) - 1
	result := make([]int, len(nums))
	i := len(nums) - 1

	for l <= r {
		left := nums[l] * nums[l]
		right := nums[r] * nums[r]
		if left < right {
			result[i] = right
			r--
		} else {
			result[i] = left
			l++
		}
		i--
	}

	return result
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}
