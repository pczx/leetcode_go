package solution24

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for curr := head; curr != nil; {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}
