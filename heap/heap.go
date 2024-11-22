package heap

import (
	"cmp"
	"errors"
)

var ErrEmptyHeap = errors.New("Empty heap")

type HeapType int8

const (
	MinHeap HeapType = iota
	MaxHeap
)

type Heap[T cmp.Ordered] struct {
	heap      []T
	compareFn func(a, b T) bool
}

func New[T cmp.Ordered](ht HeapType) *Heap[T] {
	fn := func(a, b T) bool {
		return a < b
	}

	if ht == MaxHeap {
		fn = func(a, b T) bool {
			return a > b
		}
	}

	return &Heap[T]{
		compareFn: fn,
	}
}

func (h *Heap[T]) Push(val T) {
	h.heap = append(h.heap, val)
	h.up(len(h.heap) - 1)
}

func (h *Heap[T]) Pop() (T, error) {
	if len(h.heap) == 0 {
		var zero T

		return zero, ErrEmptyHeap
	}

	out := h.heap[0]

	if len(h.heap) == 1 {
		h.heap = nil
		return out, nil
	}

	h.swap(0, len(h.heap)-1)
	h.heap = h.heap[:len(h.heap)-1]
	h.down(0)

	return out, nil
}

func (h *Heap[T]) Peek() (T, error) {
	if len(h.heap) == 0 {
		var zero T

		return zero, ErrEmptyHeap
	}

	return h.heap[0], nil
}

func (h *Heap[T]) Len() int {
	return len(h.heap)
}

func (h *Heap[T]) up(idx int) {
	if idx == 0 {
		return
	}

	pIdx := (idx - 1) >> 1

	if h.compareFn(h.heap[idx], h.heap[pIdx]) {
		h.swap(idx, pIdx)
		h.up(pIdx)
	}
}

func (h *Heap[T]) down(idx int) {
	targetIdx := idx
	leftIdx, rightIdx := (idx<<1)+1, (idx<<1)+2

	if leftIdx < len(h.heap) && h.compareFn(h.heap[leftIdx], h.heap[targetIdx]) {
		targetIdx = leftIdx
	}

	if rightIdx < len(h.heap) && h.compareFn(h.heap[rightIdx], h.heap[targetIdx]) {
		targetIdx = rightIdx
	}

	if targetIdx == idx {
		return
	}

	h.swap(targetIdx, idx)
	h.down(targetIdx)
}

func (h *Heap[T]) swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}
