/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root==nil {
		return root
	}

	if (root.Val < p.Val && root.Val >q.Val) || (root.Val > p.Val && root.Val < q.Val) {
		// split, this is the common ancestor
		return root
	}

	// else :
	if root.Val > p.Val && root.Val > q.Val {
		// check for the left subtree
		return lowestCommonAncestor(root.Left, p, q)
	}
	
	// equals case
	if p.Val==root.Val || q.Val==root.Val {
		return root
	}
	return lowestCommonAncestor(root.Right,p,q)
}
