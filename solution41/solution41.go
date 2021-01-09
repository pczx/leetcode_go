package solution41

import (
	"container/heap"
)

type Heap []int

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() (v interface{}) {
	v, *h = (*h)[h.Len()-1], (*h)[:h.Len()-1]
	return
}

func (h *Heap) Peek() (v interface{}) {
	return (*h)[0]
}

type MinHeap struct {
	Heap
}

func (h *MinHeap) Less(i, j int) bool {
	return (h.Heap)[i] > (h.Heap)[j]
}

type MaxHeap struct {
	Heap
}

func (h *MaxHeap) Less(i, j int) bool {
	return (h.Heap)[i] < (h.Heap)[j]
}

type MedianFinder struct {
	RightMin *MinHeap
	LeftMax  *MaxHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	m := new(MedianFinder)
	m.LeftMax = new(MaxHeap)
	m.RightMin = new(MinHeap)
	heap.Init(m.LeftMax)
	heap.Init(m.RightMin)
	return *m
}

func (this *MedianFinder) AddNum(num int) {
	if this.LeftMax.Len() == this.RightMin.Len() {
		heap.Push(this.RightMin, num)
		heap.Push(this.LeftMax, heap.Pop(this.RightMin))
	} else {
		heap.Push(this.LeftMax, num)
		heap.Push(this.RightMin, heap.Pop(this.LeftMax))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if (this.RightMin.Len()+this.LeftMax.Len())%2 == 0 {
		return float64(this.LeftMax.Peek().(int)+this.RightMin.Peek().(int)) / 2
	}
	return float64(this.LeftMax.Peek().(int))
}
