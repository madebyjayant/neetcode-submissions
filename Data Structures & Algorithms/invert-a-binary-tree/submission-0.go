/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	if root==nil {
		return nil
	}

	if root.Right==nil && root.Left==nil {
		return root
	}

	temp := root.Right
	root.Right = root.Left
	root.Left = temp

	_ = invertTree(root.Right)
	_ = invertTree(root.Left)

	return root
}