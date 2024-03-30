package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	result := []int{}
	l := 0
	r := len(numbers)-1

	for l < r {
		if (numbers[l] + numbers[r]) == target {
			result = append(result, l+1, r+1)
			break
		} else if (numbers[l] + numbers[r]) > target {
			r--
		} else {
			l++
		}
	}
	return result
}

func main() {
	fmt.Println(twoSum([]int{2,7,11,15}, 9)) // Output: [1,2]
	fmt.Println(twoSum([]int{2,3,4}, 6)) // Output: [1,3]
	fmt.Println(twoSum([]int{-1,0}, -1)) // Output: [1,2]
}