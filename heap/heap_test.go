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

func TestMaxHeap(t *testing.T) {
	n := 10000
	h := New(func(a, b interface{}) bool {
		x, y := a.(Data), b.(Data)
		return x.Val > y.Val
	})

	for i := 0; i < n; i++ {
		h.Push(Data{
			Val: rand.Int(),
		})

		if rand.Intn(10) > 8 {
			for j := 0; j < 3 && j < h.Size-1; j++ {
				h.Pop()
			}
		}
	}

	curr := h.Data[0].(Data).Val
	h.Pop()

	for h.Size > 0 {
		tmp := h.Pop().(Data).Val
		if tmp > curr {
			t.Fatal("Unsorted")
		}
		curr = tmp
	}
}
