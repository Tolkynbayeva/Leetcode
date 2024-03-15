package main

import "fmt"

func search(nums []int, target int) int {
	min := 0
	max := len(nums) - 1

	for min <= max {
			key := (min + max) / 2
			if target == nums[key] {
					return key
			} else if target < nums[key] {
					max = key - 1
			} else {
					min = key + 1
			}
	}
	return -1
}


func main() {
	fmt.Println(search([]int{-1,0,3,5,9,12}, 9)) //4
	fmt.Println(search([]int{-1,0,3,5,9,12}, 2)) //-1
}