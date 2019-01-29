package stack

import "fmt"

type Stack struct {
	array []int
	size  int
	top   int
}

func NewStack(size int) Stack {
	return Stack{
		array: make([]int, size),
		size:  size,
		top:   0,
	}
}

func (st *Stack) Push(item int) {
	if st.Full() == false {
		st.array[st.top] = item
		st.top++
	}
}

func (st *Stack) Pop() int {
	if st.Empty() == false {
		st.top--
		v := st.array[st.top]
		return v
	}
	return -100
}

func (st *Stack) Full() bool {
	if st.top == st.size {
		return true
	}
	return false
}

func (st *Stack) Empty() bool {
	if st.top == 0 {
		return true
	}
	return false
}

func main() {
	st := NewStack(5)
	fmt.Println(st.Empty())
	for i := 0; i < 5; i++ {
		st.Push(i)
	}
	fmt.Println(st.Full())
	fmt.Println(st.Pop())
	st.Push(6)
	for i := 0; i < 5; i++ {
		fmt.Println(st.Pop())
	}
	fmt.Println(st.Empty())
}
