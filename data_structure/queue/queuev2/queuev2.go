package queuev2

import "fmt"

type node struct {
	next *node
	val interface{}
}


type Queue struct {
	head, tail *node
	count int
}

func New() Queue {
	return Queue{}
}

func (q *Queue) Enqueue(v interface{}) {
	n := &node{nil, v}
	if q.head == nil {
		q.head, q.tail = n, n
	} else {
		q.tail.next = n
		q.tail = n
	}
	q.count++
}

func (q *Queue) Dequeue() interface{}{
	if q.head == q.tail {
		v := q.head.val
		q.head, q.tail = nil, nil
		q.count--
		return v
	}
	v := q.head.val
	q.head = q.head.next
	q.count--
	return v
}

func (q *Queue) Empty() bool {
	return q.head == nil
}

func (q *Queue) Len() int {
	return q.count
}

func main() {
	q := New()
	fmt.Println(q.Empty())
	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Empty())
}
