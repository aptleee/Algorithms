package main

import "fmt"

type heap interface {
	Less(i, j int) bool
	Swap(i, j int)
	Len() int
}

func buildHeap(h heap) {
	n := h.Len()
	for i := n/2-1; i >= 0; i-- {
		down(h, i, n)
	}
}

func up(h heap, i int) {
	for {
		j := (i-1)/2
		if j < 0 || h.Less(i, j) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

func down(h heap, i, n int) bool { // n is the length of h, so j1 can't be equal to n
	i0 := i
	smallest := i
	for {
		j1 := 2*i+1
		if j1 >= n || j1 < 0 {
			break
		}
		if h.Less(j1, smallest) {
			smallest = j1
		}
		j2 := j1+1
		if j2 < n && h.Less(j2, smallest) {
			smallest = j2
		}
		if i == smallest {
			break
		}
		h.Swap(i, smallest)
		i = smallest
	}
	return i0 < smallest
}

func Remove(h heap, i int) {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)
		if !down(h, i, n) {
			up(h, i)
		}
	}
}

func Fix(h heap, i int) {
	if !down(h, i, h.Len()) {
		up(h, i)
	}
}

func HeapSortv2(h heap) {
	buildHeap(h)
	n := h.Len()-1
	for n > 0 {
		h.Swap(0, n)
		down(h, 0, n)
		n--
	}

}

type tmp []int

func (t tmp) Len() int {
	return len(t)
}

func (t tmp) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t tmp) Less(i, j int) bool {
	return t[i] > t[j]
}

func main() {
	a := tmp{2, 3, 1, 4, 9, 6, 5}

	buildHeap(a)
	fmt.Println(a)
	HeapSortv2(a)
	fmt.Println(a)
}