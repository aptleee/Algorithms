//package binarySearchTree

//import "fmt"

package main

import "fmt"

type node struct {
	Key, Val int
	Left, Right  *node
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

func newNode(Val, Key int) *node {
	return &node{Key, Val, nil, nil}
}

func New() BST {
	return BST{}
}

func inorder(n *node) {
	if n != nil {
		inorder(n.Left)
		fmt.Println(n.Key, ":", n.Val)
		inorder(n.Right)
	}
}

func search(root *node, Key int) int {
	if root == nil {
		return -1
	}
	if root.Key > Key {
		return search(root.Left, Key)
	} else if root.Key < Key{
		return search(root.Right, Key)
	} else {
		return root.Val
	}
}


func (t *BST) Search(Key int)  {
	fmt.Println(search(t.root, Key))
}

func insert(root *node,  Key, Val int) *node {
	if root == nil {
		return newNode(Val, Key)
	}
	if Key < root.Key {
		root.Left = insert(root.Left, Key, Val)
	} else {
		root.Right = insert(root.Right, Key, Val)
	}
	return root
}

func (t *BST) Insert(Key, Val int) {
	t.root = insert(t.root, Key, Val)
}
func leftMost(root *node) *node {
	cur := root
	for cur.Left != nil {
		cur = cur.Left
	}
	return cur
}

func (t *BST) Delete(Key int) {
	t.root = bst_delete(t.root, Key)
}

func bst_delete(root *node, Key int) *node {
	if root == nil {
		return nil
	}
	if Key < root.Key {
		root.Left = bst_delete(root.Left, Key)
	} else if Key > root.Key {
		root.Right = bst_delete(root.Right, Key)
	} else {
		// this is the node to delete

		// node with one child
		if root.Left == nil {
			root = root.Right
		} else if root.Right == nil {
			root = root.Left
		} else {
			// node with two children
			d := leftMost(root.Right)
			root.Val, root.Key = d.Val, d.Key
			root.Right = bst_delete(root.Right, d.Key)
		}
	}
	return root
}

func _calculate_depth(n *node, depth int) int {
	if n == nil {
		return depth
	}
	return max(_calculate_depth(n.Left, depth+1), _calculate_depth(n.Right, depth+1))
}

func (t *BST) depth() int {
	return _calculate_depth(t.root, 0)
}

func (t *BST) Show()  {
	inorder(t.root)
}

func main() {
	t := &BST{nil}
	t.Insert(10, 1)
	t.Insert(20, 1)
	t.Insert(12, 1)
	t.Insert(22, 1)
	t.Insert(30, 1)
	t.Insert(11, 1)
	t.Insert(60, 1)
	//t.Show()
	t.Search(10)
	t.Delete(10)
	t.Show()
	t.Search(10)
	//t.root = bst_delete(t.root, 10)
	//inorder(t.root)
	//fmt.Print("\n")
	//t.root = bst_delete(t.root, 30)
	//inorder(t.root)
	//fmt.Print("\n")
	//t.root = bst_delete(t.root, 15)
	//inorder(t.root)
	//fmt.Print("\n")
	//t.root = bst_delete(t.root, 20)
	//inorder(t.root)
	//fmt.Print("\n")
	//fmt.Print(t.depth(), "\n")
}

