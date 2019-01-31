package main

import "fmt"

type node struct {
	key int
	parent, left, right *node
}

func preorder(root *node) {
	var prev *node
	for root != nil {
		if prev == root.parent { // going down, prefer left than right, maintain prev, if root has neither left
					// or right child, than go up to its parent
			fmt.Print(root.key, " ")
			prev = root
			switch {
			case root.left != nil:
				root = root.left
			case root.right != nil:
				root = root.right
			default:
				root = root.parent
			}
		} else if prev == root.left && root.right != nil { // going up from its left child, if it has right child, then go
								// down in the right direction
			prev = root
			root = root.right
		} else { // if it's going up from left child, but it doesn't have right child, just go up
			// if it's going up from right child, just going up.
			prev = root
			root = root.parent
		}
	}
}

func main() {
	n1 := &node{key:1}
	n2 := &node{key:2}
	n3 := &node{key:3}
	n4 := &node{key:4}
	n5 := &node{key:5}

	n1.left = n2
	n2.parent = n1
	n1.right = n3
	n3.parent = n1
	n2.left = n4
	n4.parent = n2
	n3.right = n5
	n5.parent = n3

	preorder(n1)

}
