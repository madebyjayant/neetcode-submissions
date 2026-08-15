/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	q := []*TreeNode{root}
	res := make([][]int, 0)
	return levelOrderTrav(q, res)
}

func levelOrderTrav(q []*TreeNode, res [][]int) [][]int {
	if len(q)==0{return res}
	temp := []int{}
	newQ := []*TreeNode{}
	for _, root := range q {
		if root==nil{
			continue
		}
		temp = append(temp, root.Val)
		newQ = append(newQ, root.Left, root.Right)
	}
	if len(temp)!=0{res = append(res, temp)}
	if len(newQ)==0{return res}
	return levelOrderTrav(newQ, res)
}
