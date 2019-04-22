package main

import (
	"Algorithms/data_structure/queue/queuev2"
	"Algorithms/data_structure/stack/stackv2"
	"fmt"
	"math"
	"strconv"
	"strings"
)


type TreeNode struct {
	Val int
	Left, Right *TreeNode
}

type Tree struct {
	root *TreeNode
}

func NewNode(v int) *TreeNode{
	return &TreeNode{v, nil,nil}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func calDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(calDepth(root.Left), calDepth(root.Right)) + 1
}

func calDepthv2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return 1 + calDepthv2(root.Right)
	}
	if root.Right == nil {
		return 1 + calDepthv2(root.Left)
	}
	return min(calDepthv2(root.Left), calDepthv2(root.Right)) + 1
}

func (t Tree) Mindepth() int {
	return calDepthv2(t.root)
}

func (t Tree) Depth() int {
	return calDepth(t.root)
}

func InOrder(root *TreeNode, res *[]int) {
	if root != nil {
		InOrder(root.Left, res)
		*res = append(*res, root.Val)
		InOrder(root.Right, res)
	}
}

func InOderv2(root *TreeNode) {
	st := stackv2.New()
	cur := root
	for cur != nil || st.Empty() == false {
		for cur != nil {
			st.Push(cur)
			cur = cur.Left
		}
		a := st.Pop().(*TreeNode)
		fmt.Print(a.Val, " ")
		cur = a.Right
	}
}

func in(root *TreeNode) {
	st := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(st) > 0 {
		for cur != nil {
			st = append(st, cur)
			cur = cur.Left
		}
		l := len(st)-1
		t := st[l]
		st = st[:l]
		fmt.Println(t.Val)
		cur = t.Right
	}
}

func PreOrder(root *TreeNode) {
	if root != nil {
		fmt.Print(root.Val, " ")
		PreOrder(root.Left)
		PreOrder(root.Right)
	}

}




func PreOrderv3(root *TreeNode) []int {
	ret := []int{}
	if root == nil {
		return ret
	}
	st := stackv2.New()
	st.Push(root)
	for st.Empty() == false {
		a := st.Pop().(*TreeNode)
		ret = append(ret, a.Val)
		if a.Right != nil {
			st.Push(a.Right)
		}
		if a.Left != nil {
			st.Push(a.Left)
		}
	}
	return ret
}

func PostOrder(root *TreeNode) {
	if root != nil {
		PostOrder(root.Left)
		PostOrder(root.Right)
		fmt.Print(root.Val, " ")
	}

}

func PostOrderv2(root *TreeNode) []int{
	ret := []int{}
	cur := root
	if cur == nil {
		return ret
	}
	st := stackv2.New()
	st.Push(root)
	for st.Empty() == false {
		r := st.Pop().(*TreeNode)
		ret = append(ret, r.Val)
		if r.Left != nil {
			st.Push(r.Left)
		}
		if r.Right != nil {
			st.Push(r.Right)
		}
	}
	Reverse(ret)
	return ret
}

func PostOrderv3(root *TreeNode) []int {
	ret := []int{}
	st := stackv2.New()
	for root != nil || st.Empty() != false {
		for root != nil {
			ret = append(ret, root.Val)
			st.Push(root)
			root = root.Right
		}
		a := st.Pop().(*TreeNode)
		root = a.Left
	}
	Reverse(ret)
	return ret
}

func LevelOrder(root *TreeNode) {
	q := queuev2.New()
	q.Enqueue(root)
	for q.Empty() == false {
		n := q.Dequeue().(*TreeNode)
		fmt.Print(n.Val, " ")
		if n.Left != nil {
			q.Enqueue(n.Left)
		}
		if n.Right != nil {
			q.Enqueue(n.Right)
		}
	}
}

func LevelOrderv2(root *TreeNode) [][]int {
	res := [][]int{}
	q := queuev2.New()
	q.Enqueue(root)
	for q.Empty() == false {
		res = append(res, []int{})
		m := q.Len()
		for i := 0; i < m; i++ {
			n := q.Dequeue().(*TreeNode)
			if n.Left != nil {
				q.Enqueue(n.Left)
			}
			if n.Right != nil {
				q.Enqueue(n.Right)
			}
			res[len(res)-1] = append(res[len(res)-1], n.Val)
		}
	}
	return res
}


func LevelOrderv3(root *TreeNode) [][]int {
	res := [][]int{}
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if  level == len(res)   {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	dfs(root, 0)
	return res
}

func ZigZagOrder(root *TreeNode) [][]int {
	res := [][]int{}
	q := queuev2.New()
	q.Enqueue(root)
	flag := 1
	for q.Empty() == false {
		res = append(res, []int{})
		m := q.Len()
		for i :=0; i < m; i++ {
			n := q.Dequeue().(*TreeNode)
			if n.Left != nil {
				q.Enqueue(n.Left)
			}
			if n.Right != nil {
				q.Enqueue(n.Right)
			}
			res[len(res)-1] = append(res[len(res)-1], n.Val)
		}
		if flag == 1 {
			flag = 0
		} else {
			Reverse(res[len(res)-1])
			flag = 1
		}
	}
	return res
}

func Reverse(a []int) {
	i, j := 0, len(a)-1
	for i < j{
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
}

func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return IsMirror(root.Left, root.Right)
}

func IsMirror(t1, t2 *TreeNode) bool {
	if t1 == nil || t2 == nil {
		return t1 == nil && t2 == nil
	}
	return t1.Val == t2.Val && IsMirror(t1.Left, t2.Right) && IsMirror(t1.Right, t2.Left)
}

func IsSymmetricv2(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return IsMirrorv2(root.Left, root.Right)
}

func IsMirrorv2(t1, t2 *TreeNode) bool{
	if t1 == nil || t2 == nil {
		return t1 == nil && t2 == nil
	}
	q := queuev2.New()
	q.Enqueue(t1)
	q.Enqueue(t2)
	for q.Empty() == false {
		t1 := q.Dequeue().(*TreeNode)
		t2 := q.Dequeue().(*TreeNode)
		if t1 == nil && t2 == nil {
			continue
		}
		if t1 == nil || t2 == nil {
			return false
		}
		if t1.Val != t2.Val {
			return false
		}
		q.Enqueue(t1.Left)
		q.Enqueue(t2.Right)
		q.Enqueue(t1.Right)
		q.Enqueue(t2.Left)
	}
	return true
}
//
func SameTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil || t2 == nil {
		return t1 == nil && t2 == nil
	}
	return t1.Val == t2.Val && SameTree(t1.Left, t2.Left) && SameTree(t1.Right, t2.Right)
}


func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return IsBalanced(root.Right) && IsBalanced(root.Left) && abs(calDepth(root.Left)-calDepth(root.Right)) <= 1
}

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}




func Construct(In, Pre []int) *TreeNode { // inorder and preorder
	var ch func(lo, hi, plo, phi int)  *TreeNode
	ch = func(lo, hi, plo, phi int) *TreeNode{
		if plo > phi || lo > hi {
			return nil
		}
		k := find(In, lo, hi, Pre[plo])
		root := NewNode(Pre[plo])
		root.Left = ch(lo,k-1,plo+1,plo+k-lo)
		root.Right = ch(k+1, hi, plo+k-lo+1, phi)
		return root
	}
	return ch(0, len(In)-1, 0, len(Pre)-1)
}

func find(A []int, lo, hi, targegt int) int {
	for j := lo; j <= hi; j++ {
		if A[j] == targegt {
			return j
		}
	}
	return -1
}

func findv2(A []int, target int) int {
	index := -1
	for i, v := range A {
		if v == target {
			index = i
			break
		}
	}
	return index
}

func buildTree(preorder []int, inorder []int) *TreeNode { // inorder and preorder
	if len(preorder) <= 0 {
		return nil
	}
	target := preorder[0]
	index := findv2(inorder, target)
	root := &TreeNode{target, nil, nil}
	root.Left = buildTree(preorder[1:index+1], inorder[:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])
	return root
}



func buildTreev3(inorder []int, postorder []int) *TreeNode { // inorder and postorder
	if len(postorder) <= 0 {
		return nil
	}
	target := postorder[len(postorder)-1]
	inx := findv2(inorder, target)
	root := &TreeNode{Val:target}
	root.Left = buildTree(inorder[:inx], postorder[:inx])
	root.Right = buildTree(inorder[inx+1:], postorder[inx:len(postorder)-1])
	return root
}





func Serializev2(root *TreeNode) string { // levelorder
	s := ""
	q := queuev2.New()
	q.Enqueue(root)
	for q.Empty() == false {
		a := q.Dequeue().(*TreeNode)
		if a == nil {
			s += "# "
		} else {
			s = s + strconv.Itoa(a.Val) + " "
			q.Enqueue(a.Left)
			q.Enqueue(a.Right)
		}
	}
	return s
}

func Deserializev2(s string) *TreeNode { //levelorder
	q := queuev2.New()
	a := strings.Split(s, " ")
	v, _ := strconv.Atoi(a[0])
	root := NewNode(v)
	q.Enqueue(root)
	i := 1
	for q.Empty() == false {
		n := q.Dequeue().(*TreeNode)
		v, _ := strconv.Atoi(a[i])
		if a[i] != "#" {
			n.Left = NewNode(v)
			q.Enqueue(n.Left)
		}
		i++
		v2, _ := strconv.Atoi(a[i])
		if a[i] != "#" {
			n.Right = NewNode(v2)
			q.Enqueue(n.Right)
		}
		i++
	}
	return root
}




func getPost(preorder []int, inorder []int) []int {
	ret := make([]int, 0)
	var helper func(preorder []int, inorder []int)
	helper = func(preorder []int, inorder []int) {
		if len(preorder) <= 0 {
			return
		}
		target := preorder[0]
		idx := findv2(inorder, target)
		helper(preorder[1:idx+1], inorder[:idx])
		helper(preorder[idx+1:], inorder[idx+1:])
		ret = append(ret, target)
	}
	helper(preorder, inorder)
	return ret
}


func Serialize(root *TreeNode) string { // preorder
	s := ""
	var sh func(root *TreeNode)
	sh = func(root *TreeNode) {
		if root == nil {
			s += "# "
		} else {
			s = s + strconv.Itoa(root.Val) + " "
			sh(root.Left)
			sh(root.Right)
		}
	}
	sh(root)
	return s
}

func Deserialize(s string) *TreeNode { // preorder
	q := queuev2.New()
	a := strings.Split(s, " ")
	for _, e := range a {
		q.Enqueue(e)
	}
	var dh func() *TreeNode
	dh = func() *TreeNode{
		if q.Empty() == true {
			return nil
		}
		a := q.Dequeue().(string)
		if a == "#" {
			return nil
		}
		t,_ := strconv.Atoi(a)
		root := NewNode(t)
		root.Left = dh()
		root.Right = dh()
		return root
	}
	return dh()
}


func LCA(root, p, q *TreeNode) *TreeNode{
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val{
		return root
	}
	left := LCA(root.Left, p, q)
	right := LCA(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

func isComplete(root *TreeNode) bool {
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for q[0] != nil {
		t := q[0]
		q = q[1:]
		q = append(q, t.Left)
		q = append(q, t.Right)
	}

	for len(q) > 0 {
		if q[0] != nil {
			return false
		}
		q = q[1:]
	}
	return true
}

func maxPathSum(root *TreeNode) int {
	maxValue := math.MinInt32

	var helper func(node *TreeNode) int
	helper = func(node *TreeNode) int{
		if node == nil {
			return 0
		}
		l := max(0, helper(node.Left))
		r := max(0, helper(node.Right))
		maxValue = max(maxValue, l + r + node.Val)
		return node.Val + max(l, r)
	}
	helper(root)
	return maxValue
}


func IsFull(root *TreeNode) bool {
	return root == nil ||
		root != nil && IsFull(root.Left) && IsFull(root.Right) && calDepth(root.Left) == calDepth(root.Right)
}

func Invert(root *TreeNode) *TreeNode{

	if root == nil {
		return nil
	}
	root.Left, root.Right = Invert(root.Right), Invert(root.Left)
	return root
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		cur.Left, cur.Right = cur.Right, cur.Left
		if cur.Left != nil {
			q = append(q, cur.Left)
		}
		if cur.Right != nil {
			q = append(q, cur.Right)
		}
	}
	return root
}

func PreOrderv2(root *TreeNode) {
	st := stackv2.New()
	cur := root
	for cur != nil || st.Empty() == false {
		for cur != nil {
			fmt.Print(cur.Val, " ")
			st.Push(cur)
			cur = cur.Left
		}
		a := st.Pop().(*TreeNode)
		cur = a.Right
	}
}