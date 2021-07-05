package main

import (
	"fmt"

	"github.com/cnnrznn/ds/stack"
)

type Hanoi struct {
	s      [3]stack.Stack
	nmoves int
}

func main() {
	for n := 1; n < 30; n++ {
		h := &Hanoi{}

		for i := n; i > 0; i-- {
			h.s[0].Push(i)
		}

		h.Solve(0, 1, 2, n)
	}
}

func (h *Hanoi) Solve(start, temp, target, n int) {
	defer func() {
		fmt.Printf("Level %v\n", n)
		fmt.Println(h.s[0])
		fmt.Println(h.s[1])
		fmt.Println(h.s[2])
		fmt.Println()
	}()

	if n == 1 {
		h.s[target].Push(h.s[start].Pop())
		h.nmoves++
		return
	}
	h.Solve(start, target, temp, n-1)
	h.s[target].Push(h.s[start].Pop())
	h.nmoves++
	h.Solve(temp, start, target, n-1)
}

func CountHanoi(n int) int {
	if n == 1 {
		return 1
	}
	return 1 + (2 * CountHanoi(n-1))
}
