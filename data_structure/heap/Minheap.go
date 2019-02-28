package main

import "fmt"

type heap interface {
	Less(i, j int) bool
	Swap(i, j int)
	Len() int
	Push(x interface{})
	Pop() interface{}
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

		if j == i || !h.Less(i, j) {
			break
		}
		h.Swap(i, j)
		i = j
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

func Push(h heap, x interface{}) {
	h.Push(x)
	up(h, h.Len()-1)
}

func Pop(h heap) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

type tmp []int

func (t tmp) Len() int {
	return len(t)
}

func (t tmp) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t tmp) Less(i, j int) bool {
	return t[i] < t[j]
}

func (t *tmp) Push(x interface{}) {
	*t = append(*t, x.(int))
}

func (t *tmp) Pop() interface{} {
	old := *t
	n := len(old)
	x := old[n-1]
	*t = old[0 : n-1]
	return x
}

func TopM(M int, input []int) *tmp{
	MinQ := &tmp{}
	buildHeap(MinQ)
	for _, x := range input {
		Push(MinQ, x)
		if MinQ.Len() > M {
			Pop(MinQ)
		}
	}
	return MinQ
}



func main() {
	input := []int{1, 3, 4, 9, 10, 5, 7, 4, 9, 10, 22, 33, 99, 1, 100, 98}

	r :=TopM(3, input)
	fmt.Println(*r)
}