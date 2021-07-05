package main

import (
	"fmt"

	"github.com/cnnrznn/ds/stack"
)

type Hanoi struct {
	s [3]stack.Stack
}

func main() {
	h := &Hanoi{}
	n := 10

	for i := n; i > 0; i-- {
		h.s[0].Push(i)
	}

	h.Step(0, 1, 2, n)
}

func (h *Hanoi) Step(start, temp, target, n int) {
	fmt.Printf("Level %v\n", n)
	fmt.Println(h.s[0])
	fmt.Println(h.s[1])
	fmt.Println(h.s[2])
	fmt.Println()
	if n == 1 {
		h.s[target].Push(h.s[start].Pop())
		return
	}
	h.Step(start, target, temp, n-1)
	h.s[target].Push(h.s[start].Pop())
	h.Step(temp, start, target, n-1)
	fmt.Printf("Level %v\n", n)
	fmt.Println(h.s[0])
	fmt.Println(h.s[1])
	fmt.Println(h.s[2])
	fmt.Println()
}
