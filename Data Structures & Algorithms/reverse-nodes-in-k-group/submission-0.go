/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	i:=0
	first := &ListNode{}
	first.Next = head
	ptr := head
	for ptr!=nil{
		if (i+1)%k==0{
			// reverse the linked list till now
			temp := ptr.Next
			ptr.Next = nil
			f, l := revLinkedList(first.Next)
			first.Next = f
			l.Next = temp
			ptr = l

			if i+1==k{
				// update the head to be this first
				head=f
			}

			first = l
		}
		ptr = ptr.Next
		i++
	}
	return head
}

func revLinkedList (head *ListNode) (f, l *ListNode) {
	first := head
	
	var prev *ListNode
	for first!=nil{
		temp := first.Next
		first.Next=prev
		prev = first
		first = temp
	}

	return prev, head
}