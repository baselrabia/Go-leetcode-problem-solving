/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func smallestFromLeaf(root *TreeNode) string {
    var helper func(*TreeNode, string) string
    
    helper = func(node *TreeNode, current string) string {
        if node == nil {
            return ""
        }
        
        ch := byte('a' + node.Val)
        current = string(ch) + current
        // Leaf node
        if node.Left == nil && node.Right == nil {
            return current
        }

        leftStr := helper(node.Left, current)
        rightStr := helper(node.Right, current)

        if leftStr == "" {
            return rightStr
        }
        if rightStr == "" {
            return leftStr
        }
        if leftStr < rightStr {
            return leftStr
        }
        return rightStr
    }
    
    return helper(root, "")
}

 