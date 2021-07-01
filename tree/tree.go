package tree

// Tree is an implementation of an AVL tree
type Tree struct {
	Root    *Node
	Compare func(a, b interface{}) int
}

type Node struct {
	Data        interface{}
	Left, Right *Node
	Height      int
}

func New(compare func(a, b interface{}) int) *Tree {
	return &Tree{
		Compare: compare,
	}
}

func (t *Tree) Insert(data interface{}) {

}
