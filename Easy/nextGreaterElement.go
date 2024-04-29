package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := []int{}
	m := make(map[int]int)
	for _, num := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < num {
			m[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}

	for i, num := range nums1 {
		if val, ok := m[num]; ok {
			nums1[i] = val
		} else {
			nums1[i] = -1
		}
	}
	return nums1
}
