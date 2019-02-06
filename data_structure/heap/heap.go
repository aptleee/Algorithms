package main

import (
	"fmt"
	"math"
)

type Heap struct {
	array    []int
	heapsize int
}

func parent(i int) int {
	return (i - 1) / 2 // 2/2 == 1   (2-1)/2 == 0
}

func child(i, j int) int {
	return i*2 + j + 1 // j == 0 or 1
}

func NewHeapA(a []int) Heap {
	return BuildHeap(a)
}

func (h *Heap) Fix(i int) {
	if h.array[i] >= h.array[parent(i)] {
		h.IncreaseKey(i, h.array[i])
	} else if h.array[i] < h.array[child(i, 0)] || h.array[i] < h.array[child(i, 1)] {
		h.MaxHeapify(i)
	} else {
		return
	}
}

func (h *Heap) MaxHeapify(i int) {
	largest := i
	for i < h.heapsize {
		left := child(i, 0)
		right := child(i, 1)
		if left < h.heapsize && h.array[largest] < h.array[left] {
			largest = left
		}
		if right < h.heapsize && h.array[largest] < h.array[right] {
			largest = right
		}
		if largest == i {
			break
		} else {
			h.array[i], h.array[largest] = h.array[largest], h.array[i]
			i = largest
		}
	}
}

func (h *Heap) MaxHeapifyv2(i int) {
	maxHeapify(h, i)
}

func  maxHeapify(h *Heap,i int) {
	largest := i
	left := child(i, 0)
	right := child(i, 1)
	if left < h.heapsize && h.array[largest] < h.array[left] {
		largest = left
	}
	if right < h.heapsize && h.array[largest] < h.array[right] {
		largest = right
	}
	if largest != i {
		h.array[i], h.array[largest] = h.array[largest], h.array[i]
		maxHeapify(h, largest)
	}
}


func (h *Heap) IncreaseKey(i, key int) {
	if h.array[i] >= key {
		fmt.Println("key must be larger than the current key")
	} else {
		for i > 0 && h.array[parent(i)] < key {
			h.array[i] = h.array[parent(i)]
			i = parent(i)
		}
		h.array[i] = key
	}
}

func (h *Heap) ExtractMax() int {
	if h.heapsize < 1 {
		fmt.Println("heap size must be larger than 0")
		return math.MaxInt64
	}
	v := h.array[0]
	h.array[0] = h.array[h.heapsize-1]
	h.heapsize--
	h.MaxHeapify(0)
	return v
}

func (h *Heap) Insert(x int) {
	h.heapsize++
	h.array[h.heapsize-1] = math.MinInt64
	h.IncreaseKey(h.heapsize-1, x)
}

func BuildHeap(a []int) Heap {
	l := len(a)
	n := Heap{
		array:    a,
		heapsize: l,
	}
	for i := l/2 - 1; i >= 0; i-- {
		n.MaxHeapifyv2(i)
	}
	return n
}

func BuildHeapv2(a []int) Heap {
	n := Heap{
		array:    a,
		heapsize: 0,
	}
	for i := 0; i < len(a); i++ {
		n.Insert(n.array[i])
	}
	return n
}

func HeapSort(a []int) {
	h := BuildHeap(a)
	for i := len(a) - 1; i > 0; i-- {
		h.array[0], h.array[i] = h.array[i], h.array[0]
		h.heapsize--
		h.MaxHeapify(0)
	}
}

func (h *Heap) Delete(i int) {
	if h.array[i] > h.array[h.heapsize-1] {
		h.array[i] = h.array[h.heapsize-1]
		h.heapsize--
		h.MaxHeapify(i)
	} else {
		h.array[i] = h.array[h.heapsize-1]
		h.IncreaseKey(i, h.array[h.heapsize-1])
		h.heapsize--
	}
}

func main() {
	a := []int{1, 6, 3, 4, 5, 8}
	// b := []int{1, 6, 3, 4, 5, 8}
	// h := BuildHeapv2(a)
	// n := BuildHeap(b)
	// h.Delete(2)
	HeapSort(a)
	fmt.Println(a)
	// fmt.Println(n)
}
