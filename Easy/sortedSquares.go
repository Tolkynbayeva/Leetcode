package main

import "fmt"

func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			nums[i] = nums[i] * -1
		}
		nums[i] = nums[i] * nums[i]
	}

	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10})) // Output: [0,1,9,16,100]
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11})) // Output: [4,9,9,49,121]
}
