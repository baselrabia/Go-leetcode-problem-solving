/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
      if root == nil {
        return 0
    }
    return sumLeavesHelper(root,0)
}

func sumLeavesHelper(node *TreeNode, rootVal int) int {
    if node == nil {
        return 0
    }
 
    rootVal = concatIntegers(rootVal, node.Val)

    // Check if it's a leaf node
    if node.Left == nil && node.Right == nil {
       return rootVal
    }

    // Recurse on both children
    return sumLeavesHelper(node.Left, rootVal) + sumLeavesHelper(node.Right, rootVal)
}

func concatIntegers(a, b int) int {
    // Convert integers to strings
    strA := strconv.Itoa(a)
    strB := strconv.Itoa(b)
    // Concatenate strings
    concatenated := strA + strB
    // Convert concatenated string back to integer
    result, _ := strconv.Atoi(concatenated)
    return result
}