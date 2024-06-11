func findTargetSumWays(nums []int, target int) int {
	var count int
	var back func(target int, index int, sum int)
	back = func(target int, index int, sum int) {
		if index == len(nums) {
			if sum == target {
				count++
			}
		} else {
			back(target, index+1, sum+nums[index])
			back(target, index+1, sum-nums[index])
		}
	}
	back(target, 0, 0)
	return count
}