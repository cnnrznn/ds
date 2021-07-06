package heap

import (
	"math/rand"
	"testing"
)

type Data struct {
	Val int
	ID  string
}

func CompareData(a, b interface{}) int {
	x := a.(Data)
	y := b.(Data)

	switch {
	case x.Val < y.Val:
		return -1
	case x.Val > y.Val:
		return 1
	default:
		return 0
	}
}

func TestHeapSort(t *testing.T) {
	n := 10
	h := New(CompareData)

	for i := 0; i < n; i++ {

		h.Push(Data{
			Val: rand.Intn(10),
		})
	}

	t.Log(h.Data)

	vals := []int{}

	for i := 0; i < n; i++ {
		vals = append(vals, h.Pop().(Data).Val)
		t.Log(h.Data[:h.Size])
	}

	for i := 1; i < n; i++ {
		if vals[i-1] > vals[i] {
			t.Fatal(vals)
		}
	}
}
