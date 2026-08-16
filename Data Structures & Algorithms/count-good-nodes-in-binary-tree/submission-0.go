/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	return finder(root, []*TreeNode{})
}

func finder(node *TreeNode, path []*TreeNode) int {
	if node==nil{return 0}
	// always append the current node
	path = append(path, node)
	// check if the node is good 
	isGood:=true
	for _,elem := range path{
		if elem.Val > node.Val{
			isGood=false
			break
		}
	}
	if isGood==true{
		if node.Left==nil && node.Right==nil{
			return 1
		}
		return 1 + finder(node.Left, path) + finder(node.Right, path)
	}else{
		return finder(node.Left, path) + finder(node.Right, path)
	}
}