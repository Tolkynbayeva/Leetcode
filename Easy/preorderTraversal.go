package main

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	var result []int
	for len(stack) > 0 {
		for i := len(stack); i > 0; i-- {
			curr := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if curr.Right != nil {
				stack = append(stack, curr.Right)
			}
			if curr.Left != nil {
				stack = append(stack, curr.Left)
			}
			result = append(result, curr.Val)
		}
	}
	return result
}
