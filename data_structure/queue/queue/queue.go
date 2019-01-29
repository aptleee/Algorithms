package queue

type Queue struct {
	array []int
	size  int
	head  int
	tail  int
}

func NewQueue(size int) Queue {
	return Queue{
		array: make([]int, size),
		size:  size,
		head:  0,
		tail:  0,
	}
}

func (q *Queue) enqueue(item int) {
	if q.Full() == false {
		q.array[q.tail] = item
		q.tail = (q.tail + 1) % q.size
	}
}

func (q *Queue) dequeue() int {
	if q.Empty() == false {
		v := q.array[q.head]
		q.head = (q.head + 1) % q.size
		return v
	}
	return -100
}

func (q *Queue) Full() bool {
	if q.head == (q.tail+1)%q.size {
		return true
	}
	return false
}

func (q *Queue) Empty() bool {
	if q.head == q.tail {
		return true
	}
	return false
}

// func main() {
// 	q := NewQueue(6)
// 	fmt.Println(q.Empty())
// 	for i := 0; i < 5; i++ {
// 		q.enqueue(i)
// 	}

// 	fmt.Println(q.Full())
// 	fmt.Println(q.dequeue())
// 	fmt.Println(q.dequeue())
// 	fmt.Println(q.Full())
// 	q.enqueue(6)
// 	fmt.Println(q.dequeue())
// 	q.enqueue(7)
// 	for i := 0; i < 5; i++ {
// 		q.dequeue()
// 	}
// 	fmt.Println(q.Empty())
// }
