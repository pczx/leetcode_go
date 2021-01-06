package solution25

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var res *ListNode

	if l1.Val < l2.Val {
		res = &ListNode{Val: l1.Val, Next: mergeTwoLists(l1.Next, l2)}
	} else {
		res = &ListNode{Val: l2.Val, Next: mergeTwoLists(l1, l2.Next)}
	}
	return res
}
