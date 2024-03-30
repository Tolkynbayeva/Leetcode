package main

import "fmt"

func removeDuplicates(nums []int) int {
	l := 0
	for r := 1; r < len(nums); r++ {
		if nums[l] != nums[r] {
			l++
			nums[l] = nums[r]
		}
	}

	nums = nums[:l+1]
	return len(nums)
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))                      // Output: 2, nums = [1,2,_]
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})) // Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
}
