package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapSort(t *testing.T) {
	n := 10
	h := New[int](func(a, b int) bool {
		return a < b
	})

	for i := 0; i < n; i++ {
		h.Push(i)
	}

	vals := []int{}

	for i := 0; i < n; i++ {
		vals = append(vals, h.Pop())
	}

	for i := 1; i < n; i++ {
		if vals[i-1] > vals[i] {
			t.Fatal(vals)
		}
	}
}

func TestMaxHeap(t *testing.T) {
	n := 10000
	h := New(func(a, b int) bool {
		return a > b
	})

	for i := 0; i < n; i++ {
		h.Push(i)
	}

	vals := []int{}

	for i := 0; i < n; i++ {
		v := h.Pop()
		vals = append(vals, v)
	}

	for i := 1; i < len(vals); i++ {
		assert.True(t, vals[i-1] > vals[i])
	}
}
