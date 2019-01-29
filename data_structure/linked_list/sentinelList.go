package main

import "fmt"

type listNode struct {
	next *listNode
	val int
}
type list struct {
	sentinel *listNode
}

func New() list{
	sen := &listNode{nil, 0}
	sen.next = sen
	return list{sen}
}

func (ll *list) Insert(v int) {
	n := &listNode{nil, v}
	m := ll.sentinel.next
	ll.sentinel.next = n
	n.next = m
}

func (ll *list) Search(v int) *listNode {
	cur := ll.sentinel.next
	for cur != ll.sentinel && cur.val != v {
		cur = cur.next
	}
	if cur == ll.sentinel {
		return nil
	}
	return cur
}

func (ll *list) Searchv2(v int) *listNode {
	cur := ll.sentinel.next
	ll.sentinel.val = v
	for cur.val != v {
		cur = cur.next
	}
	if cur == ll.sentinel {
		return nil
	}
	return cur
}

func (ll *list) Delete(v int) {
	n := ll.sentinel
	for n.next != ll.sentinel {
		if n.next.val == v {
			n.next = n.next.next
		} else {
			n = n.next
		}
	}
}

func main() {
	ll := New()
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	fmt.Println(ll.Searchv2(3).val)
	ll.Delete(2)
	fmt.Println(ll.Search(2))
}
