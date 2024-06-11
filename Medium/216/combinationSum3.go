func combinationSum3(k int, n int) [][]int {
	var result = [][]int{}
	back(&result, []int{}, 1, k, n)
	return result
}

func back(result *[][]int, combo []int, index int, k int, n int) {
	if k == 0 && n == 0 {
		comboCopy := make([]int, len(combo))
		copy(comboCopy, combo)
		*result = append(*result, comboCopy)
		return
	}

	if k < 0 || n < 0 {
		return
	}

	for i := index; i <= 9; i++ {
		combo = append(combo, i)
		back(result, combo, i+1, k-1, n-i)
		combo = combo[:len(combo)-1]
	}
}

