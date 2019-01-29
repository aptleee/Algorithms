package queue

import "fmt"

type Q struct {
	array []int
	size  int
	head  int
	tail  int
	N     int
}

func New(size int) Q {
	return Q{
		array: make([]int, size),
		size:  size,
		head:  0,
		tail:  0,
		N:     0,
	}
}

func (q *Q) enqueue(item int) {
	if q.Full() == false {
		q.array[q.tail] = item
		q.tail = (q.tail + 1) % q.size
		q.N++
	}
}

func (q *Q) dequeue() int {
	if q.Empty() == false {
		v := q.array[q.head]
		q.head = (q.head + 1) % q.size
		q.N--
		return v
	}
	return -1
}

func (q *Q) Full() bool {
	if q.N == q.size {
		return true
	}
	return false
}

func (q *Q) Empty() bool {
	if q.N == 0 {
		return true
	}
	return false
}

func main() {
	q := New(5)
	fmt.Println(q.Empty())
	for i := 0; i < 5; i++ {
		q.enqueue(i)
	}

	fmt.Println(q.Full())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.Full())
	q.enqueue(6)
	q.enqueue(7)
	for i := 0; i < 5; i++ {
		q.dequeue()
	}
	fmt.Println(q.Empty())
}
