/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	fast, slow := head, head

	if fast==nil || fast.Next==nil || fast.Next.Next==nil{
		return false
	}

	if slow==nil || slow.Next == nil {
		return false
	}

	fast = fast.Next.Next

	for fast.Next!=nil && fast.Next.Next!=nil && slow.Next!=nil{
		if fast==slow {
			return true
		}else{
			fast = fast.Next.Next
			slow = slow.Next
		}
	}

	return false
}
