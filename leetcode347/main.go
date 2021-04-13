package main

import "container/heap"

type Entry struct {
	Key, Value int
}

type EntryHeap []Entry

func (h EntryHeap) Len() int { return len(h) }

func (h EntryHeap) Less(i, j int) bool { return h[i].Value > h[j].Value }

func (h EntryHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *EntryHeap) Push(x interface{}) {
	*h = append(*h, x.(Entry))
}

func (h *EntryHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	ans := []int{}
	count := make(map[int]int)
	for _, v := range nums {
		count[v]++
	}
	var entries EntryHeap
	for k, v := range count {
		entries = append(entries, Entry{k, v})
	}
	heap.Init(&entries)
	for i := 0; i < k; i++ {
		ans = append(ans, heap.Pop(&entries).(Entry).Key)
	}
	return ans
}
