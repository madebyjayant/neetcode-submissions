/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists)==0{
		return nil
	}

	for len(lists)>1{
		mergedLists := make([]*ListNode, 0)
		for i:=0;i<len(lists);i+=2{
			l1 := lists[i]
			var l2 *ListNode
			if i+1<len(lists){
				l2 = lists[i+1]
			}
			mergedLists = append(mergedLists, mergeTwoSortedLists(l1, l2))
		}
		lists = mergedLists
	}

	return lists[0]
}

func mergeTwoSortedLists(l1, l2 *ListNode) *ListNode{
	rl := &ListNode{}
	dummy := rl

	for l1!=nil && l2!=nil {
		if l1.Val < l2.Val{
			rl.Next = l1
			l1 = l1.Next
		}else{
			rl.Next = l2
			l2=l2.Next
		}
		rl = rl.Next
	}

	if l1!=nil{
		rl.Next = l1
	}else if l2!=nil{
		rl.Next = l2
	}
	return dummy.Next
}