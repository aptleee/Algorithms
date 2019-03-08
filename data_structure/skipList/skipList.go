package main

import (
	"fmt"
	"math/rand"
	"time"
)

const maxHight = 32

type SkipListNode struct {
	value string
	next []*SkipListNode
	prev *SkipListNode
}


type SkipList struct {
	height int
	head *SkipListNode
}

func New() *SkipList {
	head := &SkipListNode{
		value:"",
		next: make([]*SkipListNode, maxHight),
	}

	return &SkipList{
		height:0,
		head:head,
	}
}

//if next().val < cur.val {
//	cur = cur.next[level]
//} else if next().val == cur.val {
// 	return cur
//} else {
//	level++
//	cur = cur.next[level]
//}
//


func (s *SkipList) Set(value string) {
	level := 0
	for ; rand.Int31n(2) == 1 && level < maxHight; level++ {
		if level > s.height {
			s.height = level
			break
		}
	}
	node := &SkipListNode{
		value:value,
		next:make([]*SkipListNode, level+1),
	}
	current := s.head

	for i := level; i >= 0; i-- {
		for ; current.next[i] != nil; current = current.next[i] {
			next := current.next[i]
			if next.value > value {
				break
			}
		}
		node.next[i] = current.next[i]
		current.next[i] = node
		node.prev = current
	}
}

func (s *SkipList) Get(value string) *SkipListNode{
	cur := s.head
	i := s.height
	for cur != nil {

		for i >= 0 {
			if cur.next[i] == nil {
				i--
			} else if cur.next[i].value < value {
				cur = cur.next[i]
				break
			} else if cur.next[i].value == value {
				return cur.next[i]
			} else {
				if i > 0 {
					i--
				}
			}
		}
	}
	return cur
}

func (s *SkipList) draw() {
	println("\n")
	ranks := make(map[string]int)
	for i, node := 0, s.head.next[0]; node != nil; node = node.next[0] {
		ranks[node.value] = i
		i++
	}

	for level := s.height; level >= 0; level-- {
		if s.head.next[level] == nil {
			continue
		}
		for i, node := 0, s.head.next[level]; node != nil; node = node.next[level] {
			rank := ranks[node.value]
			for j := 0; j < rank-i; j++ {
				print("--")
			}
			print(node.value, "-")
			i = rank + 1
		}
		println("\n")
	}
	fmt.Println("")
}

func main() {
	// change "32" to anything else to see
	rand.Seed(time.Now().Unix())
	sl := New()
	sl.Set("a")
	sl.Set("b")
	sl.Set("c")
	sl.Set("d")
	sl.Set("e")
	sl.Set("f")
	sl.Set("g")
	sl.Set("h")
	sl.Set("i")
	sl.Set("j")
	sl.Set("k")
	sl.Set("l")
	sl.Set("m")
	sl.Set("n")
	sl.Set("o")
	sl.Set("p")
	sl.Set("q")
	sl.Set("r")
	sl.Set("s")
	sl.draw()
	fmt.Println(sl.Get("q").value)
}
