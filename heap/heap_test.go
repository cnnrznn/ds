package heap

import (
	"math/rand"
	"testing"
)

type Data struct {
	Val int
	ID  string
}

func LessData(a, b interface{}) bool {
	x := a.(Data)
	y := b.(Data)

	return x.Val < y.Val
}

func TestHeapSort(t *testing.T) {
	n := 10
	h := New(LessData)

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
