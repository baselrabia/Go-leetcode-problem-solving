/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
   if depth == 1 {
        newRoot := &TreeNode{Val: val, Left: root}
        return newRoot
    }
    
    currentDepth := 1 

    dfs(root, val, depth, currentDepth+1)
    return root
}

func dfs(node *TreeNode, val, depth, currentDepth int) {
    if node == nil {
        return
    }

    if currentDepth == depth {
        // Insert nodes at the desired depth
        leftChild := &TreeNode{Val: val, Left: node.Left}
        rightChild := &TreeNode{Val: val, Right: node.Right}
        node.Left = leftChild
        node.Right = rightChild
        return
    }

    dfs(node.Left, val, depth, currentDepth+1)
    dfs(node.Right, val, depth, currentDepth+1)
}