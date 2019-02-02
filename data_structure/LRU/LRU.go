package main

import "fmt"

type doublyLinkedNode struct {
	prev, next *doublyLinkedNode
	key, val   int
}

type LRU struct {
	l, c        int
	first, last *doublyLinkedNode
	nodes       map[int]*doublyLinkedNode
}

func NewLRU(c int) LRU {
	return LRU{
		c: c,
		nodes: make(map[int]*doublyLinkedNode),
	}
}

func (lru *LRU) Get(key int) int {
	if node, ok := lru.nodes[key]; ok {
		lru.moveToFirst(node)
		return node.val
	}
	return -1
}

func (lru *LRU) Putv2(key, value int) {
	node, ok := lru.nodes[key]
	if ok {
		node.val = value
		lru.moveToFirst(node)
	} else {
		if lru.l == lru.c {
			delete(lru.nodes, lru.last.key)
			lru.removeLast()
		} else {
			lru.l++
		}
		node = &doublyLinkedNode {
			key: key,
			val: value,
		}
		lru.nodes[key] = node
		c.insertToFist(node)
	}
}

func (lru *LRU) Put(key, value int) {
	if lru.Get(key) == -1 {
		if lru.l == lru.c {
			delete(lru.nodes, lru.last.key)
			lru.removeLast()
		} else {
			lru.l++
		}
		node := &doublyLinkedNode{
			key: key,
			val: value,
		}
		lru.nodes[key] = node
		lru.insertToFirst(node)
	} else {
		lru.first.val = value
	}

}
func (lru *LRU) removeLast() {
	if lru.last.prev != nil {
		lru.last.prev.next = nil
	} else {
		lru.first = nil
	}
	lru.last = lru.last.prev
}

func (lru *LRU) moveToFirst(node *doublyLinkedNode) {
	switch node {
	case lru.first:
		return
	case lru.last:
		lru.removeLast()
	default:
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	lru.insertToFirst(node)
}

func (lru *LRU) insertToFirst(node *doublyLinkedNode) {
	if lru.last == nil {
		lru.last = node
	} else {
		node.next = lru.first
		lru.first.prev = node
	}
	lru.first = node
}

func main() {
	lru := NewLRU(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)
	fmt.Println(lru.Get(3))
	lru.Put(4, 4)
	fmt.Println(lru.Get(2))
}
