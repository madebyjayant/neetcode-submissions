/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	if root==nil{
		return true
	}
    if depth(root.Right)==depth(root.Left) || depth(root.Right)-depth(root.Left)==-1 || depth(root.Right)-depth(root.Left)==1{
		return isBalanced(root.Right) && isBalanced(root.Left)
	}
	return false
}

func depth(root *TreeNode) int {
	if root==nil {
		return 0
	}

	if root.Left==nil && root.Right==nil{
		return 1
	}

	return 1 + max(depth(root.Left),depth(root.Right))
}
