package solution40

import "container/heap"

type MaxHeap []int

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() (v interface{}) {
	v, *h = (*h)[h.Len()-1], (*h)[:h.Len()-1]
	return
}

func (h *MaxHeap) Peek() interface{} {
	return (*h)[0]
}

func getLeastNumbers(arr []int, k int) []int {
	var res []int
	n := len(arr)
	if n == 0 || k <= 0 || n < k {
		return res
	}
	pq := &MaxHeap{}
	heap.Init(pq)
	for i := 0; i < k; i++ {
		heap.Push(pq, arr[i])
	}
	for i := k; i < n; i++ {
		if pq.Len() > 0 && arr[i] < pq.Peek().(int) {
			heap.Pop(pq)
			heap.Push(pq, arr[i])
		}
	}

	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(pq).(int))
	}
	return res
}

// func getLeastNumbers(arr []int, k int) []int {
// 	if k == 0 {
// 		return []int{}
// 	}

// 	start, end := 0, len(arr)-1
// 	index := partition(arr, start, end)
// 	for index != k-1 { //一旦index==k-1，即可停止
// 		if index > k-1 {
// 			end = index - 1
// 		} else {
// 			start = index + 1
// 		}
// 		index = partition(arr, start, end)
// 	}
// 	return arr[:k]
// }

// func partition(arr []int, l, r int) int {
// 	mid := arr[l]
// 	for l < r {
// 		for l < r && arr[r] >= mid {
// 			r--
// 		}
// 		arr[l] = arr[r]
// 		for l < r && arr[l] <= mid {
// 			l++
// 		}
// 		arr[r] = arr[l]

// 	}
// 	arr[l] = mid
// 	return l
// }
