	func removeLeafNodes(root *TreeNode, target int) *TreeNode {
		if root == nil {
			return nil
		}
		root.Left = removeLeafNodes(root.Left, target)
		root.Right = removeLeafNodes(root.Right, target)
		if root.Val == target && root.Left == nil && root.Right == nil {
			return nil
		}
		return root
	}