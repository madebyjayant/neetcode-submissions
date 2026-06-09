/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var res *ListNode
	res = &ListNode{}
	ans := res
	for list1!=nil && list2!=nil{
		if list1.Val<=list2.Val{
			res.Next = list1
			res = list1
			list1 = list1.Next
		}else{
			res.Next = list2
			res = list2
			list2 = list2.Next
		}
	}

	if list1!=nil {res.Next = list1
	}else if list2!=nil{
		res.Next = list2
	}
	
	return ans.Next
}
