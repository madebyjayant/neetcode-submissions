/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head==nil {
		return head
	}
	first := head

	// create dupes
	for first!=nil{
		newNode := &Node{
			Val: first.Val,
			Next: first.Random,
		}
		first.Random = newNode
		first = first.Next
	}

	start := head
	ans := head.Random

	// fix random for the dupes
	for start!=nil {
		dupe := start.Random
		if dupe.Next!=nil{
			dupe.Random = dupe.Next.Random
		}else{
			dupe.Random = nil
		}
		start= start.Next
	}

	// fix the next values
	start = head
	for start!=nil {
		dupe := start.Random
		start.Random = dupe.Next
		if start.Next!=nil{
			dupe.Next = start.Next.Random
		}else{
			dupe.Next=nil
		}
		start = start.Next
	}
	return ans
}
