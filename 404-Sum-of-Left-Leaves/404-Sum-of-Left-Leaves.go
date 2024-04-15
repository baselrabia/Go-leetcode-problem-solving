/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {


    if root == nil {
        return 0
    }
    return sumOfLeftLeavesHelper(root.Left, true) + sumOfLeftLeavesHelper(root.Right, false)
}


func sumOfLeftLeavesHelper(node *TreeNode, isLeft bool) int {
    if node == nil {
        return 0
    }
    // Check if it's a leaf node
    if node.Left == nil && node.Right == nil {
        // If it's a left leaf, return its value
        if isLeft {
            return node.Val
        }
        return 0
    }
    // Recurse on both children
    return sumOfLeftLeavesHelper(node.Left, true) + sumOfLeftLeavesHelper(node.Right, false)
}
