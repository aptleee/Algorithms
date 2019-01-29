package stackv2
import "fmt"

type node struct {
	next *node
	val interface{}
}

type stackv2 struct {
	head *node
}

func (st *stackv2) Push(v interface{}) {
	n := &node{nil, v}
	n.next = st.head
	st.head = n
}

func (st *stackv2) Pop() {
	st.head = st.head.next
}

func (st *stackv2) Top() interface{} {
	return st.head.val
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
	fmt.Println(st.Top())
	st.Pop()
	fmt.Println(st.Top())
	st.Pop()
	fmt.Println(st.Empty())
}
