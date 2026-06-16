/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    if root==nil{
		return 0
	}

	if root.Right==nil && root.Left==nil{
		return 1
	}

	return max(1+maxDepth(root.Right), 1+maxDepth(root.Left))
} 
