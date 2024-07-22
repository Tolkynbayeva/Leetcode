/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func reverseOddLevels(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
 // var result *TreeNode
	queue := []*TreeNode{root}
	for len(queue) >0 {
		node := queue[0]
		queue = queue[1:]
		for i := len(queue); i > 0; i-- {
			if i % 2 == 1 {
				if node.Right != nil && node.Left != nil {
					node.Left.Val = node.Right.Val
					
				}
			}
		}
	}
	return root
}