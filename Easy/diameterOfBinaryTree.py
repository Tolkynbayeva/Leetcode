# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def diameterOfBinaryTree(self, root: Optional[TreeNode]) -> int:
      self.max_diametr = 0

      def depth(node):
        if node is None:
          return 0
        left = depth(node.left)
        right = depth(node.right)
        self.max_diametr = max(self.max_diametr, left+right)
        return max(left, right) + 1
      
      depth(root)
      return self.max_diametr