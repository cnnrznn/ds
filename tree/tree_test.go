package tree

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
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
}
