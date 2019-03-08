//package binarySearchTree

//import "fmt"

package main

import (
	"Algorithms/data_structure/queue/queuev2"
	"Algorithms/data_structure/stack/stackv2"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type node struct {
	Key, Val, N int
	Left, Right  *node
}

type BST struct {
	root *node
}

func (t *BST) Size() int {
	return size(t.root)
}

func size(root *node) int {
	return root.N
}
func (t *BST) Empty() bool {
	return t.Size() == 0
}

func (t *BST) Contains(Key int) bool {
	return t.Get(Key) != -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func newNode(Val, Key, N int) *node {
	return &node{Key, Val, N,nil, nil}
}

func New() BST {
	return BST{}
}

func inorder(n *node) { // zhong xu
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
func search(root *node, Key int) *node { // if not found return nil
	if root == nil {
		return nil
	}
	if root.Key > Key {
		return search(root.Left, Key)
	} else if root.Key < Key{
		return search(root.Right, Key)
	} else {
		return root
	}
}

func (t *BST) Get(Key int) int {
	if search(t.root, Key) == nil {
		return -1
	}
	return search(t.root, Key).Val
}

func (t *BST) Put(Key, Val int) {
	insert(t.root, Key, Val)
}

func insert(root *node,  Key, Val int) *node {
	if root == nil {
		return newNode(Val, Key, 1)
	}
	if Key < root.Key {
		root.Left = insert(root.Left, Key, Val)
	} else if Key > root.Key {
		root.Right = insert(root.Right, Key, Val)
	} else {
		root.Val = Val
	}
	root.N = size(root.Left) + size(root.Right) + 1
	return root
}

func leftMost(root *node) *node {
	if root == nil {
		return nil
	}
	for root.Left != nil {
		root = root.Left
	}
	return root
}

func rightMost(root *node) *node {
	if root == nil {
		return nil
	}
	for root.Right != nil {
		root = root.Right
	}
	return root
}

func (t *BST) MinKey() int {
	a := leftMost(t.root)
	if a == nil {
		return -1
	}
	return a.Key
}

func (t *BST) Floor(key int) int { // largest k less than or equal to key
	x := floor(t.root, key)
	if x == nil {
		return -1
	}
	return x.Key
}

func floor(root *node, key int) *node {
	if root == nil {
		return nil
	}
	if key == root.Key {
		return root
	}
	if key < root.Key {
		return floor(root.Left, key)
	}

	a := floor(root.Right, key)
	if a != nil {
		return a
	}
	return root
}

func ceiling(root *node, key int) *node { // if not found return nil
	if root == nil {
		return nil
	}
	if key == root.Key {
		return root
	}
	if key > root.Key {
		return floor(root.Right, key)
	}

	a := floor(root.Left, key)
	if a != nil {
		return a
	}
	return root
}
func (t *BST) Ceiling(key int) int {
	x := ceiling(t.root, key)
	if x == nil {
		return -1
	}
	return x.Key
}

func rank(root *node, key int) int {// numbers of items before key
	if root == nil {
		return 0
	}
	if root.Key < key {
		return size(root.Left) + 1 + rank(root.Right, key)
	} else if root.Key > key {
		return rank(root.Left, key)
	} else {
		return size(root.Left)
	}
}


func (t *BST) Rank(key int) int { // number of keys less than key
	return rank(t.root, key)
}

func (t *BST) Range(lo, hi int) []int {
	ret := []int{}
	var rang func(root *node, lo, hi int)
	rang = func(root *node, lo, hi int) {
		if root == nil {
			return
		}
		if lo < t.root.Key {
			rang(t.root.Left, lo, hi)
		}
		if lo <= t.root.Key && t.root.Key <= hi {
			ret = append(ret, t.root.Key)
		}
		if hi > t.root.Key {
			rang(root.Right, lo, hi)
		}
	}
	rang(t.root, lo, hi)
	return ret
}

func sel(root *node, k int) *node{
	if root == nil {
		return nil
	}
	a := size(root.Left)
	if a > k {// still the kth items in the left tree
		return sel(root.Left, k)
	} else if a < k { // the (k-a-1)th items in the right tree
		return sel(root.Right, k-a-1)
	} else {
		return root
	}

}

func (t *BST) Select(n int) int { // key == select(rank(key)) i == rank(select(i))
	if sel(t.root, n) == nil { // k is larger than size(t.root)
		return -1
	}
	return sel(t.root, n).Key
}



func (t *BST) MaxKey() int {
	a := rightMost(t.root)
	if a == nil {
		return -1
	}
	return a.Key
}

func (t *BST) Delete(Key int) {
	t.root = bst_delete(t.root, Key)
}

func deleteMin(root *node) *node {
	if root == nil {
		return nil
	}
	if root.Left == nil {
		return root.Right
	}
	root.Left = deleteMin(root.Left)
	root.N = size(root.Left) + size(root.Right) + 1
	return root
}

func bst_delete(root *node, Key int) *node { // if not found, nothing changed
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
	root.N = size(root.Left) + size(root.Right) + 1
	return root
}

func calculate_depth(n *node, depth int) int {
	if n == nil {
		return depth
	}
	return max(calculate_depth(n.Left, depth+1), calculate_depth(n.Right, depth+1))
}

func (t *BST) depth() int {
	return calculate_depth(t.root, 0)
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
				root := newNode(i, 0, 0)
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

func inorderSuccessor(root *node, p *node) *node {
	if root == nil {
		return nil
	}
	if root.Val <= p.Val {
		return inorderSuccessor(root.Right, p)
	}
	t := inorderSuccessor(root.Left, p)
	if t != nil {
		return t
	}
	return root
}

func inorderPredessor(root *node, p *node) *node {
	if root == nil {
		return nil
	}
	if root.Val >= p.Val {
		return inorderPredessor(root.Left, p)
	}
	t := inorderPredessor(root.Right, p)
	if t != nil {
		return t
	}
	return root
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
		root := newNode(0, b, 0)
		root.Left = builder(root.Key)
		root.Right = builder(maxVal)
		return root
	}
	return builder(math.MaxInt32)
}

func lowestCommonAncestor(root, p, q *node) *node {
	if root == nil {
		return nil
	}
	for {
		if (root.Val - q.Val)*(root.Val - p.Val) <= 0 {
			return root
		}
		if root.Val > q.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
}

func sumOfLeftLeaves(root *node) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return sumOfLeftLeaves(root.Right) + 1
	}
	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)

}

//func serializev2(*node) []int { //postorder
//
//}

func ValidatePreOrderOfBST(A []int) bool {
	if len(A) == 0 {
		return true
	}
	st := stackv2.New()
	low := math.MinInt32
	for _, e := range A {
		if e < low {
			return  false
		}
		for st.Empty() == false && e > st.Peek().(int) { // in the right child of one subtree
			low = st.Pop().(int)
		}
		st.Push(e)
	}
	return true
}




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

