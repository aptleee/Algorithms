package main

import "fmt"

type node struct {
	val interface{}
	next [256]*node
}

type Trie struct {
	r int
	root *node
}

func New() *Trie {
	return &Trie{
		r:256,
		root:nil,
	}
}

func search(x *node, key string, d int) *node { // d is the depth of searching
	if x == nil {
		return nil
	}
	if d == len(key) {
		return x
	}
	c := key[d]
	return search(x.next[c], key, d+1)
}

func (t *Trie) Get(key string) interface{} {
	ret := search(t.root, key, 0)
	if ret == nil {
		return nil
	}
	return ret.val
}

func insert(x *node, key string, val interface{}, d int) *node {
	if x == nil {
		x = &node{nil,[256]*node{}}
	}
	if d == len(key) {
		x.val = val
		return x
	}
	c := key[d]
	x.next[c] = insert(x.next[c], key, val, d+1)
	return x
}

func (t *Trie) Put(key string, val interface{}) {
	t.root = insert(t.root, key, val, 0)
}

func main() {
	fmt.Println([10]*node{})
}