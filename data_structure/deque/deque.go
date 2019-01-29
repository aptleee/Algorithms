package deque

import "fmt"

type deque struct {
	array []int
	N     int
	head  int
	tail  int
	size  int
}

func NewDeque(size int) deque {
	return deque{
		array: make([]int, size),
		N:     0,
		head:  0,
		tail:  size - 1,
		size:  size,
	}
}

func (dq *deque) AddFirst(item int) {
	if dq.N > dq.size {
		fmt.Println("the deque is full")
		return
	}
	dq.array[dq.head] = item
	dq.N++
	dq.head = (dq.head + 1) % dq.size
}

func (dq *deque) AddLast(item int) {
	if dq.N > dq.size {
		fmt.Println("the deque is full")
		return
	}
	dq.array[dq.tail] = item
	dq.N++
	dq.tail = (dq.tail + dq.size - 1) % dq.size
}

func (dq *deque) RemoveFirst() int {
	if dq.N <= 0 {
		fmt.Println("the deque is empty")
		return -1
	}
	dq.N--
	dq.head = (dq.head + dq.size - 1) % dq.size
	v := dq.array[dq.head]
	return v
}

func (dq *deque) RemoveLast() int {
	if dq.N <= 0 {
		fmt.Println("the deuqe is empty")
		return -1
	}
	dq.N--
	dq.tail = (dq.tail + 1) % dq.size
	v := dq.array[dq.tail]
	return v
}

func main() {
	q := NewDeque(5)
	for i := 0; i < 5; i++ {
		q.AddLast(i)
	}
	// q.RemoveFirst()
	// q.AddLast(6)
	fmt.Println(q.RemoveFirst())
	fmt.Println(q.RemoveFirst())
	q.AddFirst(6)
	q.AddFirst(7)
	fmt.Println(q.RemoveFirst())
}
