package main

import "fmt"

type MaxHeap struct {
	arr []int
}

func (h *MaxHeap) Len() int {
	return len(h.arr)
}

func (h *MaxHeap) Insert(val int) {
	h.arr = append(h.arr, val)
	h.bubbleUp(h.Len() - 1)
	h.Print()
}

func (h *MaxHeap) bubbleUp(i int) {
	if i == 0 {
		return
	}
	parent := (i - 1) / 2
	if h.arr[parent] < h.arr[i] { // swap
		h.arr[parent], h.arr[i] = h.arr[i], h.arr[parent]
		h.bubbleUp(parent)
	}
}

func (h *MaxHeap) Extract() int {
	if h.Len() == 0 {
		return 0
	}
	if h.Len() == 1 {
		return h.arr[0]
	}
	max := h.arr[0]
	h.arr[0] = h.arr[h.Len()-1]
	h.arr = h.arr[:h.Len()-1]
	h.bubbleDown(0)
	return max
}

func (h *MaxHeap) bubbleDown(i int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < h.Len() && h.arr[left] > h.arr[largest] {
		largest = left
	}
	if right < h.Len() && h.arr[right] > h.arr[largest] {
		largest = right
	}
	if largest != i {
		h.arr[i], h.arr[largest] = h.arr[largest], h.arr[i]
		h.bubbleDown(largest)
	}
}

func (h *MaxHeap) Print() {
	fmt.Println(h.arr)
}

func main() {
	h := &MaxHeap{}
	h.Insert(3)
	h.Insert(2)
	h.Insert(1)
	h.Insert(4)
	h.Insert(5)
	h.Extract()
	h.Print()
}
