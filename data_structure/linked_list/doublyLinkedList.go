package linkedlist

import (
	"fmt"
)

type doublyLinkedNode struct {
	prev, next *doublyLinkedNode
	Val        int
	Key interface{}
}

type DoublyLinkedList struct {
	head *doublyLinkedNode
}

func New(k interface{},v int) *doublyLinkedNode {
	return &doublyLinkedNode{
		Val: v,
		Key:k,
	}
}

func (dl *DoublyLinkedList) Search(k interface{}) *doublyLinkedNode{
	cur := dl.head
	for cur != nil && cur.Key != k {
		cur = cur.next
	}
	return cur
}

func (dl *DoublyLinkedList) Delete(k interface{}) {
	n := dl.Search(k)
	if n == nil {
		return
	}
	if n.prev != nil { // not the first node
		n.prev.next = n.next
	} else {// the first node
		dl.head = n.next
	}

	if n.next != nil {
		n.next.prev = n.prev
	}
}

func (dl *DoublyLinkedList) AddAtBeg(k interface{}, v int) {
	node := New(k,v)
	node.next = dl.head
	if dl.head != nil {
		dl.head.prev = node
	}
	dl.head = node
	node.prev = nil
}

func (dl *DoublyLinkedList) AddAtEnd(k interface{}, v int) {
	node := New(k,v)

	if dl.head == nil {
		dl.head = node
		return
	}
	cur := dl.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = node
	node.prev = cur
}

func (dl *DoublyLinkedList) Remove(k interface{}) {
	if dl.head == nil {
		return
	}
	ll := New(0,0)
	cur := ll
	root := dl.head
	for root != nil {
		if root.Key != k {
			cur.next = root
			root.prev = cur
			cur = cur.next
		}
		root = root.next
	}
	cur.next = nil // important when v is the last one
	dl.head = ll.next
	if dl.head.prev != nil {
		dl.head.prev = nil
	}
}



func (dl *DoublyLinkedList) DelAtBeg() int {
	if dl.head == nil {
		return -1
	}

	t := dl.head
	dl.head = t.next
	if dl.head != nil {
		dl.head.prev = nil
	}
	return t.Val
}

func (dl *DoublyLinkedList) DelAtEnd() int {
	if dl.head == nil {
		return -1
	}

	cur := dl.head
	if cur.next == nil {
		return dl.DelAtBeg()
	}

	for cur.next.next != nil {
		cur = cur.next
	}
	t := cur.next.Val
	cur.next = nil
	return t
}

func (dl *DoublyLinkedList) Count() int {
	count := 0
	for cur := dl.head; cur != nil; cur = cur.next {
		count++
	}
	return count
}

func (dl *DoublyLinkedList) Reverse() {
	cur := dl.head
	var pre, t *doublyLinkedNode
	for cur != nil {
		t = cur.next
		cur.next = pre
		cur.prev = t
		pre = cur
		cur = t
	}
	dl.head = pre
}

func (dl *DoublyLinkedList) Reversev2(p1, p2 *doublyLinkedNode) {
	pre, l:= p2, p1.prev
	var t *doublyLinkedNode
	pre.prev = p1 // if pre (p2) == nil, we can omit the statement
	for p1 != p2 {
		t = p1.next
		p1.next = pre
		p1.prev = t
		pre = p1
		p1 = t
	}
	if l == nil { // p1 is the head
		pre.prev = nil
		dl.head = pre
	} else {
		l.next = pre
		pre.prev = l
	}
}

func (dl *DoublyLinkedList) Reversev3(m, n int) {
	dummy := New(0,0)
	pre := dummy
	dummy.next = dl.head
	for i := 1; i < m; i++ {
		pre = pre.next
	}
	cur := pre.next
	var move *doublyLinkedNode
	for i := 0; i < n - m; i++ {
		move = cur.next
		cur.next = move.next
		if move.next != nil { // move.next can be nil, so it's illegal to access move.next.prev
			move.next.prev = cur
		}
		move.next = pre.next
		pre.next.prev = move
		pre.next = move
		move.prev = pre
	}
	dl.head = dummy.next
}
func (dl *DoublyLinkedList) Display() {
	for cur := dl.head; cur != nil; cur = cur.next {
		fmt.Print(cur.Val, " ")
	}
	fmt.Println()
}

func (dl *DoublyLinkedList) DisplayReverse() {
	if dl.head == nil {
		return
	}

	var cur *doublyLinkedNode
	for cur = dl.head; cur.next != nil; cur = cur.next {
	}

	for ; cur != nil; cur = cur.prev {
		fmt.Print(cur.Val, " ")
	}

	fmt.Println()
}

func Sort(dl DoublyLinkedList) DoublyLinkedList {
	if dl.head == nil || dl.head.next == nil {
		return dl
	}
	fast, slow := dl.head, dl.head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	slow.prev.next = nil
	slow.prev = nil
	return merge(Sort(dl), Sort(DoublyLinkedList{slow}))

}

func merge(l1 DoublyLinkedList, l2 DoublyLinkedList) DoublyLinkedList {
	dummy := New(0,0)
	cur := dummy
	for l1.head != nil && l2.head != nil {
		if l1.head.Val > l2.head.Val {
			cur.next = l2.head
			l2.head.prev = cur
			l2.head = l2.head.next
		} else {
			cur.next = l1.head
			l1.head.prev = cur
			l1.head = l1.head.next
		}
		cur = cur.next
	}
	if l1.head == nil {
		cur.next = l2.head
	}
	if l2.head == nil {
		cur.next = l1.head
	}
	return DoublyLinkedList{dummy.next}
}

func (dl *DoublyLinkedList) IsSorted() bool {
	if dl.head == nil || dl.head.next == nil {
		return true
	}


	for cur := dl.head; cur.next != nil; cur = cur.next {
		if cur.Val > cur.next.Val {
			return false
		}
	}
	return true
}

func (dl *DoublyLinkedList) IsCicular() bool {
	if dl.head == nil || dl.head.next == nil {
		return false
	}
	slow, fast := dl.head, dl.head.next
	for slow != fast {
		if fast == nil || fast.next == nil {
			return false
		}
		slow = slow.next
		fast = fast.next.next
	}
	return true
}

//func main() {
//	ll := DoublyLinkedList{}
//	ll.AddAtBeg(10, 10)
//	ll.AddAtEnd(20,20)
//	ll.AddAtBeg(20,20)
//	ll.AddAtBeg(30,20)
//	ll.AddAtEnd(40,40)
//	ll.AddAtBeg(50,40)
//	ll.Display()
//	fmt.Println("sorted? ", ll.IsSorted())
//	fmt.Println("has cycle?", ll.IsCicular())
//	//fmt.Println(ll.head.prev)
//	//ll.Reverse()
//	//ll.Display()
//	ll.Remove(20)
//	ll.Display()
//	//ll.Reversev2(ll.head.next, ll.head.next.next.next)
//	ll.Reversev3(1, 4)
//	ll.Display()
//	ll = Sort(ll)
//	ll.Display()
//	fmt.Println("sorted? ",ll.IsSorted())
//	//cur := ll.head
//	//for ; cur.next != nil; cur = cur.next {
//	//}
//	//cur.next = ll.head
//	//ll.head.prev = cur
//	//fmt.Println("has cycle?", ll.IsCicular())
//	// ll.DisplayReverse()
//	// fmt.Println(ll.DelAtBeg())
//	// fmt.Println(ll.DelAtEnd())
//	// fmt.Println("Display")
//	// ll.Display()
//	// fmt.Println(ll.DelAtBeg())
//	//ll.Display()
//	// fmt.Print(ll.DelAtBeg())
//	// ll.Display()
//}
