package algorithm_test

import (
	"algorithm"
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	h := algorithm.Heap[int]{}
	h.Push(1)
	h.Push(2)
	h.Push(3)
	h.Push(4)
	h.Push(5)
	h.Push(6)
	h.Push(7)
	h.Push(8)
	h.Push(9)
	fmt.Println(h.Pop(), h.Pop(), h.Pop(), h.Pop(), h.Pop(), h.Pop(), h.Pop(), h.Pop(), h.Pop())
}
