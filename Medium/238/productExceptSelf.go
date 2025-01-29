package main

import "fmt"

func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	answer[0] = 1
	for i := 1; i < len(nums); i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] *= right
		right *= nums[i]
	}

	return answer
}

func main() {
	nums := []int{1,2,3,4}
	fmt.Println(productExceptSelf(nums))
}