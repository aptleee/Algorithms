package d_ray_heap

import "OS_and_Networks_in_golang/aptlee/swap"
import "fmt"
import "math"

func parent(i, d int) int {
	return ((i - 1) / d)
}

func child(i, d, c int) int {
	return (d*i + c + 1)
}

type D_ray_heap struct {
	elements  []int
	d         int
	heapSize int
}

func Max_heapify(h *D_ray_heap, i int) {
	largest := i
	basechild := child(i, h.d, 0)
	for k := 0; k < h.d; k++ {
		child := basechild + k
		if child < h.heapSize && h.elements[child] > h.elements[largest] {
			largest = child
		}
	}

	if largest != i {
		swap.Swap(h.elements, i, largest)
		Max_heapify(h, largest)
	}
}

func Max_heapify_iter(h *D_ray_heap, i int) {
	for i < h.heapSize {
		largest := i
		basechild := child(i, h.d, 0)
		for k := 0; k < h.d; k++ {
			child := basechild + k
			if child < h.heapSize && h.elements[child] > h.elements[largest] {
				largest = child
			}
		}

		if largest != i {
			swap.Swap(h.elements, i, largest)
			i = largest
		} else {
			break
		}
	}
}

func Extract_max(h *D_ray_heap) int {
	if h.heapSize < 1 {
		fmt.Println("heap size must be larger than 0")
		return math.MinInt64
	}
	max := h.elements[0]
	h.elements[0] = h.elements[h.heapSize-1]
	h.heapSize--
	Max_heapify(h, 0)
	return max
}

func Increase_key(h *D_ray_heap, i, key int) {
	if h.elements[i] >= key {
		fmt.Println("key must be larger than the current key")
	} else {
		for i > 0 && h.elements[parent(i, h.d)] < key {
			h.elements[i] = h.elements[parent(i, h.d)]
			i = parent(i, h.d)
		}
		h.elements[i] = key
	}
}

func Insert(h *D_ray_heap, key int) {
	h.heapSize++
	h.elements[h.heapSize-1] = math.MinInt64
	Increase_key(h, h.heapSize-1, key)
}
