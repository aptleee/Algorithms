//package binarySearchTree

//import "fmt"

package main

import "fmt"

type node struct {
	key   int
	val   int
	left  *node
	right *node
}

type BST struct {
	root *node
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func newNode(val, key int) *node {
	return &node{key, val, nil, nil}
}

func inorder(n *node) {
	if n != nil {
		inorder(n.left)
		fmt.Print(n.val, " ")
		inorder(n.right)
	}
}

func Search(root *node, key int) int {
	if root == nil {
		return -1
	}
	if root.key < key {
		return Search(root.left, key)
	} else if root.key > key{
		return Search(root.right, key)
	} else {
		return root.val
	}
}
func insert(root *node, val, key int) *node {
	if root == nil {
		return newNode(val, key)
	}
	if key < root.key {
		root.left = insert(root.left, val, key)
	} else {
		root.right = insert(root.right, val, key)
	}
	return root
}

func inorderSuccessor(root *node) *node {
	cur := root
	for cur.left != nil {
		cur = cur.left
	}
	return cur
}

func bst_delete(root *node, key int) *node {
	if root == nil {
		return nil
	}
	if key < root.key {
		root.left = bst_delete(root.left, key)
	} else if key > root.key {
		root.right = bst_delete(root.right, key)
	} else {
		// this is the node to delete

		// node with one child
		if root.left == nil {
			root = root.right
		} else if root.right == nil {
			root = root.left
		} else {
			// node with two children
			d := inorderSuccessor(root.right)
			root.val, root.key = d.val, d.key
			root.right = bst_delete(root.right, d.key)
		}
	}
	return root
}

// helper function for t.depth
func _calculate_depth(n *node, depth int) int {
	if n == nil {
		return depth
	}
	return max(_calculate_depth(n.left, depth+1), _calculate_depth(n.right, depth+1))
}

func (t *BST) depth() int {
	return _calculate_depth(t.root, 0)
}


func main() {
	t := &BST{nil}
	inorder(t.root)
	t.root = insert(t.root, 10)
	t.root = insert(t.root, 20)
	t.root = insert(t.root, 15)
	t.root = insert(t.root, 30)
	fmt.Print(t.depth(), "\n")
	inorder(t.root)
	fmt.Print("\n")
	t.root = bst_delete(t.root, 10)
	inorder(t.root)
	fmt.Print("\n")
	t.root = bst_delete(t.root, 30)
	inorder(t.root)
	fmt.Print("\n")
	t.root = bst_delete(t.root, 15)
	inorder(t.root)
	fmt.Print("\n")
	t.root = bst_delete(t.root, 20)
	inorder(t.root)
	fmt.Print("\n")
	fmt.Print(t.depth(), "\n")
}

