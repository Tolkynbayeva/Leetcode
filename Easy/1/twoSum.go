package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if index, found := numMap[complement]; found {
			return []int{index, i}
		}
		numMap[num] = i
	}
  return []int{} 
}

func main() {
	nums := []int{2,7,11,15}
	target := 9
	fmt.Println(twoSum(nums, target))
}