// Package tree implements a balanced (AVL) binary search tree
package tree

import (
	"fmt"
)

type Tree[T any] interface {
	Insert(x T)
	Find(x T) *T
	Height() int
}

type avlTree[T any] struct {
	Root    *Node[T]
	Compare func(a, b T) int // < 0 if a < b, > 0 if a > b, 0 iff a == b
}

type Node[T any] struct {
	Left, Right *Node[T]
	Height      int
	Val         T
}

func New[T any](Compare func(a, b T) int) Tree[T] {
	return &avlTree[T]{
		Compare: Compare,
	}
}

func (t *avlTree[T]) Height() int {
	return t.Root.height()
}

func (t *avlTree[T]) Find(x T) *T {
	u := t.Root
	for u != nil {
		if t.Compare(x, u.Val) == 0 {
			return &u.Val
		} else if t.Compare(x, u.Val) < 0 {
			u = u.Left
		} else {
			u = u.Right
		}
	}
	return nil
}

func (t *avlTree[T]) Insert(Val T) {
	t.Root = t.insertLeaf(t.Root, Val)
}

func (t *avlTree[T]) insertLeaf(u *Node[T], Val T) *Node[T] {
	if u == nil {
		return &Node[T]{
			Val:    Val,
			Height: 1,
		}
	}

	if t.Compare(Val, u.Val) < 0 { // insert left
		u.Left = t.insertLeaf(u.Left, Val)
		u.Height = max(u.Left.height(), u.Right.height()) + 1

		if u.Balance() == -2 {
			if u.Left.Balance() > 0 { // left-right
				u = t.rotateLeftRight(u)
			} else { // right
				u = t.rotateRight(u)
			}
		}
	} else { // insert right
		u.Right = t.insertLeaf(u.Right, Val)
		u.Height = max(u.Left.height(), u.Right.height()) + 1

		if u.Balance() == 2 {
			if u.Right.Balance() < 0 {
				u = t.rotateRightLeft(u)
			} else {
				u = t.rotateLeft(u)
			}
		}
	}

	return u
}

func (t *avlTree[T]) rotateLeft(u *Node[T]) *Node[T] {
	x, z := u, u.Right
	t2 := z.Left

	x.Right = t2
	z.Left = x

	x.Height = max(x.Left.height(), x.Right.height()) + 1
	z.Height = max(z.Left.height(), z.Right.height()) + 1

	return z
}

func (t *avlTree[T]) rotateRight(u *Node[T]) *Node[T] {
	x, z := u, u.Left
	t2 := z.Right

	x.Left = t2
	z.Right = x

	x.Height = max(x.Left.height(), x.Right.height()) + 1
	z.Height = max(z.Left.height(), z.Right.height()) + 1

	return z
}

func (t *avlTree[T]) rotateLeftRight(u *Node[T]) *Node[T] {
	x, z, y := u, u.Left, u.Left.Right
	t2 := y.Left
	t3 := y.Right

	z.Right = t2
	y.Left = z

	x.Left = t3
	y.Right = x

	x.Height = max(x.Left.height(), x.Right.height()) + 1
	z.Height = max(z.Left.height(), z.Right.height()) + 1
	y.Height = max(y.Left.height(), y.Right.height()) + 1

	return y
}

func (t *avlTree[T]) rotateRightLeft(u *Node[T]) *Node[T] {
	x, z, y := u, u.Right, u.Right.Left
	t2 := y.Left
	t3 := y.Right

	z.Left = t3
	y.Right = z

	x.Right = t2
	y.Left = x

	x.Height = max(x.Left.height(), x.Right.height()) + 1
	z.Height = max(z.Left.height(), z.Right.height()) + 1
	y.Height = max(y.Left.height(), y.Right.height()) + 1

	return y
}

func (u *Node[T]) Balance() int {
	return u.Right.height() - u.Left.height()
}

func (u *Node[T]) height() int {
	if u == nil {
		return 0
	}
	return u.Height
}

func (t *avlTree[T]) String() string {
	return t.buildTreeString(t.Root, "", 0)
}

func (t *avlTree[T]) buildTreeString(u *Node[T], prefix string, level int) string {
	if u == nil {
		return ""
	}

	levelStr := ""
	for _ = range level {
		levelStr += " |"
	}
	levelStr += " "

	result := ""

	result += fmt.Sprintf("%v %v (%v, %v)\n", levelStr, prefix, u.Val, u.Height)
	result += t.buildTreeString(u.Left, "L ", level+1)
	result += t.buildTreeString(u.Right, "R ", level+1)

	return result
}
