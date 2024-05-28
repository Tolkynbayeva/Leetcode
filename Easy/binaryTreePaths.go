func binaryTreePaths(root *TreeNode) []string {
	var paths []string
	backtrack(root, "", &paths)
	return paths
}

func backtrack(node *TreeNode, path string, paths *[]string) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		*paths = append(*paths, path+strconv.Itoa(node.Val))
		return
	}
	if node.Left != nil {
		backtrack(node.Left, path+strconv.Itoa(node.Val)+"->", paths)
	}

	if node.Right != nil {
		backtrack(node.Right, path+strconv.Itoa(node.Val)+"->", paths)
	}
}