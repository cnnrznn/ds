// Package tree implements a balanced (AVL) binary search tree
package tree

import (
	"fmt"
)

type Tree[T any] struct {
	Root *Node[T]
	Less func(a, b T) bool
}

type Node[T any] struct {
	Left, Right *Node[T]
	Height      int
	Val         T
}

func New[T any](Less func(a, b T) bool) *Tree[T] {
	return &Tree[T]{
		Less: Less,
	}
}

func (t *Tree[T]) Insert(Val T) {
	t.Root, _ = t.insertLeaf(t.Root, Val)
}

func (t *Tree[T]) insertLeaf(u *Node[T], Val T) (*Node[T], bool) {
	if u == nil {
		return &Node[T]{
			Val: Val,
		}, true
	}

	var changed bool

	if t.Less(Val, u.Val) {
		u.Left, changed = t.insertLeaf(u.Left, Val)
		if !changed {
			return u, false
		}

		if t.balance[u] < 0 { // doubly left heavy
			if t.balance[u.Left] > 0 { // left-right
				u = t.rotateLeftRight(u)
			} else {
				u = t.rotateRight(u)
			}
		} else {
			t.balance[u] -= 1
			if t.balance[u] == 0 {
				return u, false
			}
		}
	} else {
		u.Right, changed = t.insertLeaf(u.Right, Val)
		if !changed {
			return u, false
		}

		if t.balance[u] > 0 { // doubly right heavy
			if t.balance[u.Right] < 0 { // right-left
				u = t.rotateRightLeft(u)
			} else { // left
				u = t.rotateLeft(u)
			}
		} else { // left heavy or neutral
			t.balance[u] += 1
			if t.balance[u] == 0 {
				return u, false
			}
		}
	}

	return u, true
}

func (t *Tree[T]) rotateLeft(u *Node[T]) *Node[T] {
	x, z := u, u.Right
	t2 := z.Left

	x.Right = t2
	z.Left = x

	if t.balance[z] == 0 {
		t.balance[x] = 1
		t.balance[z] = -1
	} else {
		t.balance[x] = 0
		t.balance[z] = 0
	}

	return z
}

func (t *Tree[T]) rotateRight(u *Node[T]) *Node[T] {
	x, z := u, u.Left
	t2 := z.Right

	x.Left = t2
	z.Right = x

	if t.balance[z] == 0 {
		t.balance[x] = -1
		t.balance[z] = 1
	} else {
		t.balance[x] = 0
		t.balance[z] = 0
	}

	return z
}

func (t *Tree[T]) rotateLeftRight(u *Node[T]) *Node[T] {
	x, z, y := u, u.Left, u.Left.Right
	t2 := y.Left
	t3 := y.Right

	z.Right = t2
	y.Left = z

	x.Left = t3
	y.Right = x

	if t.balance[y] == 0 {
		t.balance[x] = 0
		t.balance[z] = 0
	} else if t.balance[y] < 0 {
		t.balance[x] = 1
		t.balance[z] = 0
	} else {
		t.balance[x] = 0
		t.balance[z] = -1
	}

	t.balance[y] = 0

	return y
}

func (t *Tree[T]) rotateRightLeft(u *Node[T]) *Node[T] {
	x, z, y := u, u.Right, u.Right.Left
	t2 := y.Left
	t3 := y.Right

	z.Left = t3
	y.Right = z

	x.Right = t2
	y.Left = x

	if t.balance[y] == 0 { // height t2 same as t3
		t.balance[x] = 0
		t.balance[z] = 0
	} else if t.balance[y] < 0 { // t2 higher
		t.balance[x] = 0
		t.balance[z] = 1
	} else { // t3 higher
		t.balance[x] = -1
		t.balance[z] = 0
	}

	t.balance[y] = 0

	return y
}

func (t *Tree[T]) String() string {
	return t.buildTreeString(t.Root, "", 0)
}

func (t *Tree[T]) buildTreeString(u *Node[T], prefix string, level int) string {
	if u == nil {
		return ""
	}

	levelStr := ""
	for _ = range level {
		levelStr += " |"
	}
	levelStr += " "

	result := ""

	result += fmt.Sprintf("%v %v %v\n", levelStr, prefix, u.Val)
	result += t.buildTreeString(u.Left, "L ", level+1)
	result += t.buildTreeString(u.Right, "R ", level+1)

	return result
}

func (t *Tree[T]) PrintBalance() {
	for k, v := range t.balance {
		fmt.Printf("%v: %v\n", k.Val, v)
	}
}
