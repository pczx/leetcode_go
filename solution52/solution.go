package solution52

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	lenA, lenB := getLength(headA), getLength(headB)
	gap := int(math.Abs(float64(lenA - lenB)))
	if lenA > lenB {
		for i := 0; i < gap; i++ {
			curB = curB.Next
		}
	} else {
		for i := 0; i < gap; i++ {
			curA = curA.Next
		}
	}
	for ; curA != nil && curB != nil; curA, curB = curA.Next, curB.Next {
		if curA == curB {
			return curA
		}
	}
	return nil
}

func getLength(head *ListNode) int {
	l := 0
	for ; head != nil; head = head.Next {
		l++
	}
	return l
}