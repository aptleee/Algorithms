package linkedlist

import "fmt"

type singlyLinkedNode struct {
	next *singlyLinkedNode
	val  int
}

type singlyLinkedList struct {
	head *singlyLinkedNode
}

func NewNode(v int) *singlyLinkedNode {
	return &singlyLinkedNode{nil, v}
}


func (sl *singlyLinkedList) Search(v int) *singlyLinkedNode{
	cur := sl.head
	for cur != nil && cur.val != v {
		cur = cur.next
	}
	return cur
}

func (sl *singlyLinkedList) AddAtBeg(v int) {
	node := NewNode(v)
	node.next = sl.head
	sl.head = node
}

func (sl *singlyLinkedList) AddAtEnd(v int) {
	node := NewNode(v)

	if sl.head == nil {
		sl.head = node
		return
	}
	cur := sl.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = node
}

func (sl *singlyLinkedList) Remove(v int) {
	if sl.head == nil {
		return
	}
	ll := NewNode(0)
	cur := ll
	for sl.head != nil {
		if sl.head.val != v {
			cur.next = sl.head
			cur = cur.next
		}
		sl.head = sl.head.next
	}
	cur.next = nil // important when v is the last one
	sl.head = ll.next
}



func (sl *singlyLinkedList) DelAtBeg() int {
	if sl.head == nil {
		return -1
	}

	t := sl.head
	sl.head = t.next
	return t.val
}

func (sl *singlyLinkedList) DelAtEnd() int {
	if sl.head == nil {
		return -1
	}
	cur := sl.head
	if cur.next == nil {
		sl.head = nil
		return cur.val
	}
	for cur.next.next != nil {
		cur = cur.next
	}
	t := cur.next.val
	cur.next = nil
	return t
}

func (sl *singlyLinkedList) Count() int {
	count := 0
	for cur := sl.head; cur != nil; cur = cur.next {
		count++
	}
	return count
}

func (sl *singlyLinkedList) Delete(v int) {
	n := sl.Search(v)
	if n == nil {
		fmt.Printf("not found %d", v)
		fmt.Println()
		return
	}
	if n.next != nil { // n is not the last node
		n.val = n.next.val
		n.next = n.next.next
	} else { // only one node in sl
		if n == sl.head {
			sl.head = nil
		} else { // find the node previous to the last node
			var cur *singlyLinkedNode
			for cur = sl.head; cur.next.next != nil; cur = cur.next {}
			cur.next = nil
		}
	}
}

func (sl *singlyLinkedList) Deletev2(v int) {
	dummy := NewNode(0)
	dummy.next = sl.head
	cur := dummy
	for cur.next != nil {
		if cur.next.val == v {
			cur.next = cur.next.next
		} else {
			cur = cur.next
		}
	}
	sl.head = dummy.next
}
func (sl *singlyLinkedList) Reverse() {
	cur := sl.head
	var pre, t *singlyLinkedNode
	for cur != nil {
		t = cur.next
		cur.next = pre
		pre = cur
		cur = t
	}
	sl.head = pre
}

func (sl *singlyLinkedList) Reversev2(p2 *singlyLinkedNode) {
	cur, pre := sl.head, p2
	var t *singlyLinkedNode
	for cur != p2 {
		t = cur.next
		cur.next = pre
		pre = cur
		cur = t
	}
	sl.head = pre
}

func (sl *singlyLinkedList) Reversev3(m, n int) {
	dummy := NewNode(0)
	pre := dummy
	pre.next = sl.head
	for i := 1; i < m; i++ {
		pre = pre.next
	}
	cur := pre.next
	var move *singlyLinkedNode
	for i := 0; i < n-m; i++ {
		move = cur.next
		cur.next = move.next
		move.next = pre.next
		pre.next = move
	}
	sl.head = dummy.next
}

func (sl *singlyLinkedList) Display() {
	for cur := sl.head; cur != nil; cur = cur.next {
		fmt.Print(cur.val, " ")
	}
	fmt.Println()
}

func sort(sl singlyLinkedList) singlyLinkedList {
	if sl.head == nil || sl.head.next == nil {
		return sl
	}
	fast, slow, pre := sl.head, sl.head, sl.head
	for fast != nil && fast.next != nil {
		pre = slow
		slow = slow.next
		fast = fast.next.next
	}
	pre.next = nil
	return smerge(sort(sl), sort(singlyLinkedList{slow}))

}

func smerge(l1 singlyLinkedList, l2 singlyLinkedList) singlyLinkedList {
	dummy := NewNode(0)
	cur := dummy
	for l1.head != nil && l2.head != nil {
		if l1.head.val > l2.head.val {
			cur.next = l2.head
			l2.head = l2.head.next
		} else {
			cur.next = l1.head
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
	return singlyLinkedList{dummy.next}
}

func (sl *singlyLinkedList) IsSorted() bool {
	if sl.head == nil || sl.head.next == nil {
		return true
	}
	var cur *singlyLinkedNode
	for cur = sl.head; cur.next != nil; cur = cur.next {
		if cur.val > cur.next.val {
			return false
		}
	}
	return true
}

func (sl *singlyLinkedList) IsCicular() bool {
	if sl.head == nil || sl.head.next == nil {
		return true
	}
	slow, fast := sl.head, sl.head.next

	for slow != fast {
		if fast == nil || fast.next == nil {
			return false
		}
		slow = slow.next
		fast = fast.next.next
	}
	return true
}

func detectCyclev2(head *singlyLinkedNode) *singlyLinkedNode {
	if head == nil || head.next == nil {
		return nil
	}

	us := make(map[*singlyLinkedNode]struct{})
	cur := head
	for cur != nil {
		if _, ok := us[cur]; ok {
			return cur
		}
		us[cur] = struct{}{}
		cur = cur.next
	}
	return cur
}

func detectCycle(head *singlyLinkedNode) *singlyLinkedNode {
	if head == nil || head.next == nil {
		return nil
	}
	entry, slow, fast := head, head, head
	for fast == nil || fast.next == nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			for slow != entry {
				slow = slow.next
				entry = entry.next
			}
			return entry
		}
	}
	return nil
}




func smergev2(h1 *singlyLinkedNode, h2 *singlyLinkedNode) *singlyLinkedNode {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}
	if h1.val < h2.val {
		h1.next = smergev2(h1.next, h2)
		return h1
	}
	h2.next = smergev2(h2.next, h1)
	return h2
}
//func main() {
//	ll := singlyLinkedList{}
//
//	ll.AddAtBeg(10)
//	//ll.AddAtEnd(20)
//	//ll.AddAtBeg(20)
//	//ll.AddAtBeg(30)
//	ll.AddAtEnd(40)
//	//ll.AddAtBeg(50)
//	//ll.Display()
//	//ll.Delete(20)
//	//ll.Display()
//	ll.Deletev2(40)
//	ll.Deletev2(10)
//	ll.Display()
//	//fmt.Println("sorted? ", ll.IsSorted())
//	//fmt.Println("has cycle?", ll.IsCicular())
//	////fmt.Println(ll.head.prev)
//	////ll.Reverse()
//	////ll.Display()
//	//ll.Remove(20)
//	//ll.Display()
//	////ll.Reversev2(ll.head.next, ll.head.next.next.next)
//	//ll.Reversev3(1, 4)
//	//ll.Display()
//	//ll = sort(ll)
//	//ll.Display()
//	//fmt.Println("sorted? ",ll.IsSorted())
//
//}
