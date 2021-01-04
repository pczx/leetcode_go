package solution18

type ListNode struct {
	Val  int
	Next *ListNode
}

// func deleteNode(head *ListNode, val int) *ListNode {
// 	if head == nil {
// 		return nil
// 	}

// 	if head.Val == val {
// 		return head.Next
// 	}

// 	for curr := head; curr.Next != nil; curr = curr.Next {
// 		if curr.Next.Val == val {
// 			curr.Next = curr.Next.Next
// 			break
// 		}
// 	}

// 	return head
// }

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Val: -1}
	dummy.Next = head
	pre, curr := dummy, head
	for curr != nil {
		if curr.Val == val {
			pre.Next = curr.Next
			break
		}
		pre = curr
		curr = curr.Next
	}
	return dummy.Next
}
