# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def smallestFromLeaf(self, root: Optional[TreeNode]) -> str:
        def helper(node: TreeNode, current_path: str) -> str:
            if not node:
                return ""

            # Convert the node value to its corresponding character
            ch = chr(ord('a') + node.val)

            # Concatenate the current character to the current path
            current_path += ch

            # If the node is a leaf, return the current path
            if node.left is None and node.right is None:
                return current_path[::-1]  # Reverse the current path to get the string from leaf to root

            # Recursively get the smallest string from left and right subtrees
            left_str = helper(node.left, current_path)
            right_str = helper(node.right, current_path)

            # Choose the lexicographically smallest string
            if not left_str:
                return right_str
            if not right_str:
                return left_str
            if left_str < right_str:
                return left_str
            return right_str
        return helper(root, "")
    