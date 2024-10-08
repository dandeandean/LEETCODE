package main

//import "fmt"

type MaxHeap []int

func (h MaxHeap) Len() int      { return len(h) }
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Pop() int {
	if h.Len() == 0 {
		panic("WHAT ARE YOU DOING CALLING POP ON AN EMPTY HEAP?")
	}
	old := *h
	x := old[0]
	if h.Len() <= 1 {
		*h = []int{}
	} else {
		*h = old[1:]
	}
	for i := len(*h) / 2; i >= 0; i-- {
		h.MaxHeapify(i)
	}
	return x
}
func (h *MaxHeap) Push(n int) {
	*h = append(*h, n)
	for i := len(*h) / 2; i >= 0; i-- {
		h.MaxHeapify(i)
	}
	//fmt.Println(h)
}

func HeapFrom(nums []int) *MaxHeap {
	heap := new(MaxHeap)
	for _, n := range nums {
		heap.Push(n)
	}
	return heap
}

func (h MaxHeap) MaxHeapify(i int) {
	left, right := i*2+1, i*2+2
	ibig := i
	if left < len(h) {
		if h[left] >= h[ibig] {
			ibig = left
		}
	}
	if right < len(h) {
		if h[right] >= h[ibig] {
			ibig = right
		}
	}
	if ibig != i {
		h.Swap(ibig, i)
		h.MaxHeapify(ibig)
	}
}
