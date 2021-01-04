package solution22

type ListNode struct {
	Val  int
	Next *ListNode
}

// func getKthFromEnd(head *ListNode, k int) *ListNode {

// 	var getK func(*ListNode) *ListNode

// 	getK = func(head *ListNode) *ListNode {
// 		if head == nil || k == 0 {
// 			return nil
// 		}

// 		res := getK(head.Next)
// 		k--
// 		if k == 0 {
// 			return head
// 		}
// 		return res

// 	}
// 	return getK(head)
// }

func getKthFromEnd(head *ListNode, k int) *ListNode {

	if head == nil || k == 0 {
		return nil
	}

	fast, slow := head, head

	for i := 0; i < k-1; i++ {
		if fast.Next != nil {
			fast = fast.Next
		} else {
			return nil
		}
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
