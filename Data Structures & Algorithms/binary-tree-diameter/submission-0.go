/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	if root==nil {
		return 0
	}

	
	longestPath := depth(root.Left) + depth(root.Right)

	leftLongestPath := diameterOfBinaryTree(root.Left)
	rightLongestPath := diameterOfBinaryTree(root.Right)

	return max(longestPath, max(leftLongestPath, rightLongestPath))
}

func depth(root *TreeNode) int {
	if root==nil{
		return 0
	}

	if root.Left==nil && root.Right==nil{
		return 1
	}

	return 1+max(depth(root.Left), depth(root.Right))
}