package main

import "fmt"

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	maxAmount := 0
	for l < r {
		currAmount := min(height[l], height[r]) * (r - l)
		if currAmount > maxAmount {
			maxAmount = currAmount
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxAmount
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
