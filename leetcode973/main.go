package main

import "container/heap"

type PHeap [][]int

func (h PHeap) Len() int { return len(h) }

func (h PHeap) Less(i, j int) bool {
	return h[i][0]*h[i][0]+h[i][1]*h[i][1] < h[j][0]*h[j][0]+h[j][1]*h[j][1]
}

func (h PHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *PHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *PHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	ans := [][]int{}
	v := PHeap(points)
	heap.Init(&v)
	for i := 0; i < k; i++ {
		ans = append(ans, heap.Pop(&v).([]int))
	}
	return ans
}
