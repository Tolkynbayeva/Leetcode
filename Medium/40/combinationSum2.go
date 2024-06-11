func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result [][]int
	var back func(combi []int, target int, start int)
	back = func(combi []int, target int, start int) {
		if target == 0 {
			combiCopy := make([]int, len(combi))
			copy(combiCopy, combi)
			result = append(result, combiCopy)
		}
		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			if candidates[i] > target {
				break
			}
			combi = append(combi, candidates[i])
			back(combi, target-candidates[i], i+1)
			combi = combi[:len(combi)-1]
		}
	}
	back([]int{}, target, 0)
	return result
}