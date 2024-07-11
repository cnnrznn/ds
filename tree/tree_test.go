package tree

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

// TODO use breadth-first traversal to assert tree structure after operations

func TestInsert(t *testing.T) {
	t.Run("123 tree", func(t *testing.T) {
		a := New(func(a, b int) bool {
			return a < b
		})

		a.Insert(1)
		a.Insert(2)
		a.Insert(3)

		fmt.Println(a)
	})

	t.Run("10 tree", func(t *testing.T) {
		a := New(func(a, b int) bool {
			return a < b
		})

		for i := 1; i <= 10; i++ {
			a.Insert(i)
			fmt.Println(a)
			fmt.Println("=====================")
		}
	})

	t.Run("6 random", func(t *testing.T) {
		a := New(func(a, b int) bool { return a < b })

		for _, n := range rand.Perm(6) {
			a.Insert(n)
		}

		fmt.Println(a)
	})

	t.Run("100 height", func(t *testing.T) {
		a := New[int](func(a, b int) bool {
			return a > b
		})

		for _, n := range rand.Perm(10000) {
			a.Insert(n)
		}

		fmt.Printf("%v: %v\n", a.Root.Height, math.Log2(float64(10000)))
	})
}
