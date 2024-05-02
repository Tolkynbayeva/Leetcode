package main

func maxWidthRamp(nums []int) int {
	stack := []int{}
	maxWidth := 0

	for i := 0; i < len(nums); i++ {
		if len(stack) == 0 || nums[i] <= nums[stack[len(stack)-1]] {
			stack = append(stack, i)
			// fmt.Println("st", stack)
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[i] {
			width := i - stack[len(stack)-1]
			if width > maxWidth {
				maxWidth = width
			}
			stack = stack[:len(stack)-1]
		}

	}
	return maxWidth
}
