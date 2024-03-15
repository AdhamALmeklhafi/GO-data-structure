package main

import "fmt"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// will find the max heap from botton to top which calld heapify
func (h *MaxHeap) maxHeapifyUp(index int) {

	// if the parent heap is smaller than the index than the swap function will work
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)

		// update the loop after swaping
		index = parent(index)
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChildIndex(i int) int {
	return 2*i + 1
}

func rightChildIndex(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]

}

func main() {
	maxheap := &MaxHeap{}

	// in the front it will be the largest key always
	buildHeap := []int{80, 100, 50}
	for _, v := range buildHeap {
		maxheap.Insert(v)

		// the output will be 100,80,50
		fmt.Println(maxheap)
	}
}
