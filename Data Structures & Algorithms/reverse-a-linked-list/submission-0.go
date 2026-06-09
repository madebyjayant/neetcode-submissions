/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head==nil || head.Next==nil{
		return head
	}
	var prev *ListNode
	for head.Next!=nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	head.Next = prev
	return head
}
