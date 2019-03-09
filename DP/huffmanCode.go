package main

type TreeNode struct {
	Left, Right *TreeNode
	Val string
	fre float64
}

func Huffman(A []*TreeNode) *TreeNode{
	PQ := Init(A)
	n := len(A)
	for i := 0; i < n-1; i++ {
		node := new(TreeNode)
		t1 := Pop(PQ)
		t2 := Pop(PQ)
		node.fre = t1.fre + t2.fre
		Push(PQ, node)
	}
	return Pop(PQ)
}