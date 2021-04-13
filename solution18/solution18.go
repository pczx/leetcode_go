package solution18

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{-1,head}

	for cur := dummy; cur.Next != nil; cur = cur.Next{
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		}
	}
	return dummy.Next
}

func deleteNode2(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	if head.Val == val {
		return head.Next
	}
	head.Next = deleteNode2(head.Next, val)
	return head
}