package main

import (
	"AlgorithmsGo-master/data_structure/stack/stackv2"
	"fmt"
)


type TreeNode struct {
	val int
	left, right *TreeNode
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

func calDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(calDepth(root.left), calDepth(root.right)) + 1
}

func (t Tree) Depth() int {
	return calDepth(t.root)
}

func InOrder(root *TreeNode) {
	if root == nil {
		return
	}
	InOrder(root.left)
	fmt.Print(root.val, " ")
	InOrder(root.right)
}

func InOderv2(root *TreeNode) {
	st := stackv2.New()
	var a *TreeNode
	cur := root
	for cur != nil || st.Empty() == false {
		//PushAllLeft(cur, st)
		for cur != nil {
			st.Push(cur)
			cur = cur.left
		}
		a = st.Top().(*TreeNode)
		st.Pop()
		fmt.Print(a.val, " ")
		cur = a.right
	}
}

//func PushAllLeft(root *TreeNode, st []*TreeNode) {
//	for root != nil {
//		st = append(st, root)
//		root = root.left
//	}
//}

func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.val, " ")
	PreOrder(root.left)
	PreOrder(root.right)
}

//func PreOderv2(root *TreeNode) {
//
//}
//
//func PostOrderv2(root *TreeNode) {
//
//}

func PostOrder(root *TreeNode) {
	if root == nil {
		return
	}
	PostOrder(root.left)
	PostOrder(root.right)
	fmt.Print(root.val, " ")
}

func LevelOrder(root *TreeNode) {
	var q []*TreeNode // queue
	var n *TreeNode   // temporary node

	q = append(q, root)

	for len(q) != 0 {
		n, q = q[0], q[1:]
		fmt.Print(n.val, " ")
		if n.left != nil {
			q = append(q, n.left)
		}
		if n.right != nil {
			q = append(q, n.right)
		}
	}
}


//func ZigZagOrder(root *TreeNode) {
//
//}
//func IsSymmetric(root *TreeNode) bool {
//
//}
//
//func IsMirror(t1 *TreeNode, t2 *TreeNode) bool {
//
//}
//
//func SameTree(t1 *TreeNode, t2 *TreeNode) bool {
//
//}
//
//
//
//
//func MinDepth(root *TreeNode) int {
//
//}
//
//func IsBalanced(root *TreeNode) bool {
//
//}
//
//func Invert(root *TreeNode) {
//
//}
//

func main() {


	t := Tree{}
	t.root = NewNode(0)
	t.root.left = NewNode(1)
	t.root.right = NewNode(2)
	t.root.left.left = NewNode(3)
	t.root.left.right = NewNode(4)
	t.root.right.left = NewNode(5)
	t.root.right.right = NewNode(6)
	t.root.right.right.right = NewNode(10)

	InOrder(t.root)
	fmt.Println()
	InOderv2(t.root)

	//PreOrder(t.root)
	//fmt.Print("\n")
	//PostOrder(t.root)
	//fmt.Print("\n")
	//LevelOrder(t.root)
	//fmt.Print("\n")
	//fmt.Print(t.Depth(), "\n")

}
