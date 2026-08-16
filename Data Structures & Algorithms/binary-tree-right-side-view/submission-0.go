/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	q := []*TreeNode{root}
	res := make([]int, 0)
	if root==nil{return res}
	return right(q, res)
}

func right(q []*TreeNode, res []int) []int {
	if len(q)!=0{
		rightMost := q[len(q)-1]
		if rightMost!=nil{
			res = append(res, q[len(q)-1].Val)
		}
	}

	newQ := make([]*TreeNode, 0)
	for _, root := range q {
		if root.Left!=nil{
			newQ = append(newQ, root.Left)
		}
		if root.Right!=nil{
			newQ = append(newQ, root.Right)
		}
	}
	if len(newQ)==0{return res}

	return right(newQ, res)
}
