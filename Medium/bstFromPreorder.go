func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	stack := []*TreeNode{root}
	for i := 1; i < len(preorder); i++ {
		node := stack[len(stack)-1]
		child := &TreeNode{Val: preorder[i]}
		for len(stack) > 0 && stack[len(stack)-1].Val < child.Val {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		if node.Val < child.Val {
			node.Right = child
		} else {
			node.Left = child
		}
		stack = append(stack, child)
	}
	return root
}
