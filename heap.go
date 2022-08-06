package algorithm

import "fmt"

type Heap[T int | int8 | int16 | int32 | int64 | float64 | float32 | string] struct {
	slice []T
}

func (h *Heap[T]) Pop() (result T) {
	if len(h.slice) == 0 {
		return result
	}
	result = h.slice[0]
	h.slice[0] = h.slice[len(h.slice)-1]
	//h.slice[len(h.slice)-1] = nil
	h.slice = h.slice[:len(h.slice)-1]
	fmt.Println(h)
	h.down(0)
	fmt.Println(h)
	return result
}

func (h *Heap[T]) Push(value T) {

	h.slice = append(h.slice, value)
	h.up(len(h.slice) - 1)
	//fmt.Println(h)
}

func (h *Heap[T]) up(index int) {
	if index == 0 {
		return
	}
	for i := (index - 1) / 2; i >= 0 && h.slice[i] < h.slice[index]; i, index = (i-1)/2, i {
		h.slice[i], h.slice[index] = h.slice[index], h.slice[i]
	}
}

func (h *Heap[T]) down(index int) {
	for i := index; i < len(h.slice); {
		maxIndex := 0
		if len(h.slice) > i+2 {
			if h.slice[i+1] > h.slice[i+2] {
				maxIndex = i + 1
			} else {
				maxIndex = i + 2
			}
		} else if len(h.slice) == i+2 {
			maxIndex = i + 1
		} else {
			return
		}

		if h.slice[i] < h.slice[maxIndex] {
			h.slice[i], h.slice[maxIndex] = h.slice[maxIndex], h.slice[i]
			i = maxIndex
		} else {
			return
		}
	}
}
