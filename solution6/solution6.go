package solution6

// ListNode defination
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	res := []int{}
	for pNode := head; pNode != nil; pNode = pNode.Next {
		res = append(res, pNode.Val)
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
