package main

func minSubArrayLen(target int, nums []int) int {
	start := 0
	currSum := 0
	minLen := len(nums) + 1
	for end := 0; end < len(nums); end++ {
		currSum += nums[end]
		for currSum >= target {
			if end-start+1 < minLen {
				minLen = end - start + 1
			}
			currSum -= nums[start]
			start++
		}
	}
	if minLen == len(nums)+1 {
		return 0
	}
	return minLen
}
