package solution35

// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	cloneNodes(head)
	connectRandomNodes(head)
	return reconnectNode(head)
}

func cloneNodes(head *Node) {
	for node := head; node != nil; {
		cloned := &Node{Val: node.Val}
		cloned.Next = node.Next
		node.Next = cloned
		node = cloned.Next
	}
}

func connectRandomNodes(head *Node) {
	for node := head; node != nil; {
		cloned := node.Next
		if node.Random != nil {
			cloned.Random = node.Random.Next
		}
		node = cloned.Next
	}
}

func reconnectNode(head *Node) *Node {
	node := head
	var cloned *Node
	var clonedHead *Node
	if node != nil {
		cloned = node.Next
		clonedHead = cloned
		node.Next = cloned.Next
		node = node.Next
	}

	for node != nil {
		cloned.Next = node.Next
		cloned = cloned.Next
		node.Next = cloned.Next
		node = node.Next
	}

	return clonedHead
}
