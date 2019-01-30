package main

import (
	"AlgorithmsGo-master/data_structure/queue/queuev2"
	"AlgorithmsGo-master/data_structure/stack/stackv2"
	"fmt"
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


func PreOrder(root *TreeNode) {
	if root != nil {
		fmt.Print(root.Val, " ")
		PreOrder(root.Left)
		PreOrder(root.Right)
	}

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

func PostOrderv2(root *TreeNode) []int{
	ret := []int{}
	st := stackv2.New()
	cur := root
	if cur == nil {
		return ret
	}
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

func PostOrder(root *TreeNode) {
	if root != nil {
		PostOrder(root.Left)
		PostOrder(root.Right)
		fmt.Print(root.Val, " ")
	}

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
	return IsMirrorv2(root, root)
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

func Invert(root *TreeNode) *TreeNode{

	if root == nil {
		return nil
	}
	root.Left, root.Right = Invert(root.Right), Invert(root.Left)
	return root
}

func Invertv2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	q := queuev2.New()
	q.Enqueue(root)
	for q.Empty() == false {
		root = q.Dequeue().(*TreeNode)
		root.Left, root.Right = root.Right, root.Left
		if root.Left != nil {
			q.Enqueue(root.Left)
		}
		if root.Right != nil {
			q.Enqueue(root.Right)
		}
	}
	return root
}

//func Construct(A []int) []*TreeNode {
//
//}
//
func Serialize(root *TreeNode) string { // preorder
	s := ""
	sh(root, &s)
	return s
}

func sh(root *TreeNode, s *string) {
	if root == nil {
		*s += "# "
	} else {
		*s = *s + strconv.Itoa(root.Val) + " "
		sh(root.Left, s)
		sh(root.Right, s)
	}
}
func Deserialize(s string) *TreeNode {
	q := queuev2.New()
	a := strings.Split(s, " ")
	for _, e := range a {
		q.Enqueue(e)
	}

	return dh(q)
}

func dh(q *queuev2.Queue) *TreeNode {
	a := q.Dequeue().(string)
	if a == "#" {
		return nil
	}
	t,_ := strconv.Atoi(a)
	root := NewNode(t)
	root.Left = dh(q)
	root.Right = dh(q)
	return root
}

//func Serializev2(root *TreeNode) string { // levelorder
//
//}
//

func Copy(root *TreeNode) *TreeNode{
	if root == nil {
		return nil
	}
	r2 := NewNode(root.Val)
	r2.Left = Copy(root.Left)
	r2.Right = Copy(root.Right)
	return r2
}




func main() {
	t := Tree{}
	t.root = NewNode(0)
	t.root.Left = NewNode(1)
	t.root.Right = NewNode(2)
	t.root.Left.Left = NewNode(3)
	t.root.Left.Right = NewNode(4)
	t.root.Right.Left = NewNode(5)
	t.root.Right.Right = NewNode(6)
	t.root.Right.Right.Right = NewNode(10)
	fmt.Println(Serialize(t.root))
	a := Deserialize(Serialize(t.root))

	PreOrder(a)
	//fmt.Println("root1 is symmetric", IsSymmetric(root1))
	fmt.Println("is SameTree", SameTree(t.root, a))
	//Invertv2(root1)
	//fmt.Println("root1 is balanced", IsBalanced(root1))
	//fmt.Println("is mirror", IsMirrorv2(root1, t.root) )
	//fmt.Println("is SameTree", SameTree(t.root, root1))
	////fmt.Println()
	////PreOrder(t.root)
	////fmt.Println()
	////PreOrderv2(t.root)
	////fmt.Println()
	//fmt.Println(LevelOrderv2(t.root))
	//fmt.Println(LevelOrderv3(root1))
	////fmt.Println(ZigZagOrder(t.root))

	//fmt.Print("\n")
	//fmt.Println(t.Depth())
	//fmt.Println(t.Mindepth())
	//LevelOrder(t.root)
	//fmt.Println()
	//Invertv2(t.root)
	//LevelOrder(t.root)
	//fmt.Println()
}
