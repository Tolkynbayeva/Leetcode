func permute(nums []int) [][]int {
	var result [][]int
	backtrack(nums, 0, &result)
	return result
}

func backtrack(nums []int, pos int, result *[][]int) {
	if pos == len(nums) {
		perm := make([]int, len(nums))
		copy(perm, nums)
		*result = append(*result, perm)
		return
	}

	for i := pos; i < len(nums); i++ {
		nums[i], nums[pos] = nums[pos], nums[i]
		backtrack(nums, pos+1, result)
		nums[i], nums[pos] = nums[pos], nums[i]
	}
}