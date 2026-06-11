/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	slow, fast := head, head.Next


	if head==nil || head.Next==nil{
		return
	}

	// find the middle point
	for fast!=nil && fast.Next!=nil{
		slow = slow.Next
		fast = fast.Next.Next
	}

	// end the first list
	second := slow.Next
	slow.Next=nil

	// reverse the second list
	var prev *ListNode
	for second!=nil {
		temp := second.Next
		second.Next = prev
		prev = second
		second = temp
	}
	
	first := head
	second = prev

	// merge 
	for second!=nil {
		temp1 := first.Next
		first.Next = second
		temp2 := second.Next
		second.Next = temp1
		second = temp2
		first = temp1
	}

	return
}
