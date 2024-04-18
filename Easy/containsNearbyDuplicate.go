package main

import "fmt"

func containsNearbyDuplicate(nums []int, n int) bool {
	n = n + 1
	start := 0
	for i := 0; i < len(nums); i++ {
		start = i
		for k := start + 1; k < i+n && k < len(nums); k++ {
			if nums[start] == nums[k] {
				return true
			}
		}
		start++
	}

	return false
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))       // true
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))       // true
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2)) // false
}
