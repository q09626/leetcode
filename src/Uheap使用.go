package main

import (
	"container/heap"
	"fmt"
)

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func main() {
	h := &intHeap{2, 1, 5, 6, 4, 3, 7, 9, 8, 0}
	heap.Init(h)
	fmt.Println(*h)

	fmt.Println(heap.Pop(h).(int))
	fmt.Println(heap.Pop(h).(int))
	fmt.Println(heap.Pop(h).(int))

	heap.Push(h, 6)

	fmt.Println(*h)

	for len(*h) > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
