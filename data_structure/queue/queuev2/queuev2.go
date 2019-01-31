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

func New() *Queue {
	return &Queue{}
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
	v := q.head.val
	q.count--
	if q.head == q.tail {
		q.head, q.tail = nil, nil
		return v
	}
	q.head = q.head.next
	return v
}

func (q *Queue) Empty() bool {
	return q.head == nil
}

func (q *Queue) Len() int {
	return q.count
}

func (q *Queue) Front() interface{} {
	return q.head.val
}

func (q *Queue) Back() interface{} {
	return q.tail.val
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
