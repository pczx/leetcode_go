package main

/**
 * lru design
 * @param operators int整型二维数组 the ops
 * @param k int整型 the k
 * @return int整型一维数组
 */

type LRUCache struct {
	Head     *Node
	Tail     *Node
	Capacity int
	Cache    map[int]*Node
}

type Node struct {
	Key   int
	Value int
	Prev  *Node
	Next  *Node
}

func Construct(capacity int) *LRUCache {
	head, tail := new(Node), new(Node)
	head.Prev = tail
	head.Next = tail
	tail.Next = head
	tail.Prev = head
	return &LRUCache{head, tail, capacity, make(map[int]*Node)}
}

func (lru *LRUCache) set(key, value int) {
	node, ok := lru.Cache[key]
	
	if ok {
		node.Value = value
	} else {
		node := &Node{key, value, nil, nil}
		lru.Cache[key] = node
		lru.addToHead(node)
		if len(lru.Cache) > lru.Capacity {
			tail := lru.removeTail()
			delete(lru.Cache, tail.Key)
		}
	}
}

func (lru *LRUCache) get(key int) int {
	node, ok := lru.Cache[key]
	if !ok {
		return -1
	}
	lru.moveToHead(node)
	return node.Value
}

func (lru *LRUCache) moveToHead(node *Node) {
	lru.removeNode(node)
	lru.addToHead(node)
}

func (lru *LRUCache) removeNode(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (lru *LRUCache) addToHead(node *Node) {
	node.Next = lru.Head.Next
	node.Prev = lru.Head
	lru.Head.Next.Prev = node
	lru.Head.Next = node
}

func (lru *LRUCache) removeTail() *Node {
	node := lru.Tail.Prev
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	return node
}

func LRU(operators [][]int, k int) []int {
	// write code here
	ans := []int{}
	lru := Construct(k)
	for _, op := range operators {
		if op[0] == 1 {
			lru.set(op[1], op[2])
		} else if op[0] == 2 {
			ans = append(ans, lru.get(op[1]))
		}
	}
	return ans
}

func main() {
	LRU([][]int{{1, 1, 1}, {1, 2, 2}, {1, 3, 2}, {2, 1}, {1, 4, 4}, {2, 2}}, 3)
}
