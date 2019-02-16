package Tree

// 2-node, 3-node and 4-node
// red link lean left
// no node connected by two red links
// the tree has perfect black balance

type color bool

const (
	RED  color = true
	BLACK color = false
)

type Key interface {
	Less(i Key) bool
}

type TreeNode struct {
	Key Key
	Val interface{}
	Color color
	N int
	Left, Right *TreeNode
}

type RBT struct {
	root *TreeNode
}

func size(node *TreeNode) int {
	return node.N
}

func RotateLeft(node *TreeNode) *TreeNode{
	x := node.Right
	node.Right = x.Left
	x.Left = node
	x.Color = node.Color
	node.Color = RED
	x.N = node.N
	node.N = 1 + size(node.Left) + size(node.Right)
	return x
}



func RotateRight(node *TreeNode) *TreeNode{
	x := node.Left
	node.Left = x.Right
	x.Right = node
	x.Color = node.Color
	node.Color = RED
	x.N = node.N
	node.N = 1 + size(node.Left) + size(node.Right)
	return x
}

func flipColor(node *TreeNode) {
	node.Left.Color, node.Right.Color = BLACK, BLACK
	node.Color = RED
}

func isRed(node *TreeNode) bool {
	return node.Color == RED
}

func search(key Key, node *TreeNode) *TreeNode{
	if node == nil {
		return nil
	}
	if key.Less(node.Key){
		return search(key, node.Left)
	} else if !key.Less(node.Key) {
		return search(key, node.Right)
	}
	return node
}

func insert(key Key, val interface{}, node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	if key.Less(node.Key) {
		node.Left = insert(key, val, node.Left)
	} else if !key.Less(node.Key) {
		node.Right = insert(key, val, node.Right)
	} else {
		node.Val = val
	}

	// sequence counts, insert to a 2-node or insert to a 3-node
	if isRed(node.Right) && !isRed(node.Left) {
		RotateLeft(node)
	} else if isRed(node.Left) && isRed(node.Left.Left) {
		RotateRight(node)
	} else if isRed(node.Left) && isRed(node.Right) { // decompose 4-node
		flipColor(node)
	}

	node.N = 1 + size(node.Left) + size(node.Right)

	return node
}

// can only remove 3-node, if it's 2-node, we need to transform
func moveRedLeft(node *TreeNode) {

}



func (T *RBT) Get(key Key) interface{} {
	n := search(key, T.root)
	if n == nil {
		return nil
	}
	return n.Val
}

func (T *RBT) Put(key Key, val interface{}) {
	T.root = insert(key, val, T.root)
	T.root.Color = BLACK
}