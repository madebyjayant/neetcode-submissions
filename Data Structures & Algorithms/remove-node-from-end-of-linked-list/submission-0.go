/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	total := 0
	first, ans := head, head

	// find the total count
	for head!=nil{
		head=head.Next
		total++
	}

	// now walk till the node before the target
	count := 0
	target := total-n

	if target==0 {
		return first.Next
	}
	
	for first!=nil {
		if count==target-1{
			t := first.Next
			first.Next = t.Next
			break
		}
		first=first.Next
		count++
	}
	return ans
}
