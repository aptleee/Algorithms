//package binarySearchTree

//import "fmt"

package main

import (
	"AlgorithmsGo-master/data_structure/queue/queuev2"
	"fmt"
	"math"
	"strconv"
	"strings"
)

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

func preOrder(n *node) {
	if n != nil {
		fmt.Print(n.Key, " ")
		preOrder(n.Left)
		preOrder(n.Right)
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

func UniqueBST(n int) int { // n denotes the number of nodes
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i < n+1; i++ {
		for j := 1; j <= i; j++ { // either dp[j-1] or dp[i-j] can be dp[0], so the border condition is [1, i]
			dp[i] += dp[j-1] * dp[i-j]// [0, j-1] + [j+1, i]
		}
	}
	return dp[n]
}

func UniqueBSTv2(n int) []*node {
	// two intervals
	return generating(1, n)
}

func generating(start, end int) []*node {
	res := []*node{}
	if start > end {
		res = append(res, nil)
		return res
	}

	for i := start; i <= end; i++ { // take care of the border condition
		l := generating(start, i-1)
		r := generating(i+1, end)
		for j := 0; j < len(l); j++ {
			for k := 0; k < len(r); k++ {
				root := newNode(i, 0)
				root.Left = l[j]
				root.Right = r[k]
				res = append(res, root)
			}
		}
	}
	return res
}

func levelorder(root *node) {
	if root != nil {
		q := queuev2.New()
		q.Enqueue(root)
		for q.Empty() == false {
			a := q.Dequeue().(*node)
			fmt.Print(a.Val, " ")
			if a.Left != nil {
				q.Enqueue(a.Left)
			}
			if a.Right != nil {
				q.Enqueue(a.Right)
			}
		}
	}

}


//func deserializev2(post string) *node { // postorder
//
//	var build func(post string, root *node)
//	build = func(post string, root *node) {
//
//	}
//
//}


func serialize(n *node) string { // preorder
	s := ""
	var sh func(n *node)
	sh = func(n *node){
		if n != nil {
			s = s + strconv.Itoa(n.Key) + " "
			sh(n.Left)
			sh(n.Right)
		}
	}
	sh(n)
	return s[:len(s)-1]
}

func deserialize(pre string) *node{ // preorder
	q := queuev2.New()
	for _, e := range strings.Split(pre, " ") {
		v, _ := strconv.Atoi(e)
		q.Enqueue(v)
	}
	var builder func(maxVal int) *node
	builder = func(maxVal int) *node {
		if q.Empty() == true || q.Front().(int) >= maxVal {
			return nil
		}
		b := q.Dequeue().(int)
		root := newNode(0, b)
		root.Left = builder(root.Key)
		root.Right = builder(maxVal)
		return root
	}
	return builder(math.MaxInt32)
}



//func serializev2(*node) []int { //postorder
//
//}





func main() {
	//t := &BST{nil}
	//t.Insert(25, 1)
	//t.Insert(20, 1)
	//t.Insert(12, 1)
	//t.Insert(22, 1)
	//t.Insert(30, 1)
	//t.Insert(11, 1)
	//t.Insert(60, 1)
	//fmt.Println(serialize(t.root))
	//a := deserialize(serialize(t.root))
	//preOrder(a)
	//fmt.Println()
	//preOrder(t.root)
	s := "s s"
	fmt.Println(s[:3])
	//t.Search(10)
	//t.Delete(10)
	//t.Show()
	//t.Search(10)
	//fmt.Println(UniqueBST(3))
	//for _, e := range UniqueBSTv2(3) {
	//	levelorder(e)
	//	fmt.Println()
	//}
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

