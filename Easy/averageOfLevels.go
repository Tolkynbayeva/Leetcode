	func averageOfLevels(root *TreeNode) []float64 {
		if root == nil {
			return nil
		}
		var result []float64
		queue := []*TreeNode{root}
		for len(queue) > 0 {
			var level []int
			for i := len(queue); i > 0; i-- {
				node := queue[0]
				queue = queue[1:]
				level = append(level, node.Val)
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			}
			var sum float64
			for _, v := range level {
				sum += float64(v)
			}
			result = append(result, sum/float64(len(level)))
		}
		return result
	}