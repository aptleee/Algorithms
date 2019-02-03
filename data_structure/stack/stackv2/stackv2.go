package stackv2
import "fmt"

type node struct {
	next *node
	Val interface{}
}

type stackv2 struct {
	head *node
}

func (st *stackv2) Push(v interface{}) {
	n := &node{nil, v}
	n.next = st.head
	st.head = n
}

func (st *stackv2) Pop() interface{}{
	t := st.head.Val
	st.head = st.head.next
	return t
}

func (st *stackv2) Peek() interface{} {
	return st.head.Val
}
func (st *stackv2) Empty() bool {
	return st.head == nil
}

func New() stackv2{
	return stackv2{}
}

func main() {
	st := New()
	st.Push(true)
	st.Push(false)
	fmt.Println(st.Peek())
	st.Pop()
	fmt.Println(st.Peek())
	st.Pop()
	fmt.Println(st.Empty())
}
