package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeap(t *testing.T) {
	h := New[int](MaxHeap)

	t.Run("Push", func(t *testing.T) {
		h.Push(10)
		h.Push(20)
		h.Push(5)

		assert.Equal(t, 3, h.Len())
		assert.Equal(t, 20, h.heap[0])
	})

	t.Run("Peek", func(t *testing.T) {
		h.Push(15)
		top, err := h.Peek()

		assert.NoError(t, err)
		assert.Equal(t, 20, top)
	})

	t.Run("Pop", func(t *testing.T) {
		val, err := h.Pop()

		assert.NoError(t, err)
		assert.Equal(t, 20, val)

		val, err = h.Pop()

		assert.NoError(t, err)
		assert.Equal(t, 15, val)

		assert.Equal(t, 2, h.Len())
	})
}

func TestMinHeap(t *testing.T) {
	h := New[int](MinHeap)

	t.Run("Push", func(t *testing.T) {
		h.Push(10)
		h.Push(20)
		h.Push(5)

		assert.Equal(t, 3, h.Len())
		assert.Equal(t, 5, h.heap[0])
	})

	t.Run("Peek", func(t *testing.T) {
		h.Push(3)
		top, err := h.Peek()

		assert.NoError(t, err)
		assert.Equal(t, 3, top)
	})

	t.Run("Pop", func(t *testing.T) {
		val, err := h.Pop()

		assert.NoError(t, err)
		assert.Equal(t, 3, val)

		val, err = h.Pop()

		assert.NoError(t, err)
		assert.Equal(t, 5, val)

		assert.Equal(t, 2, h.Len())
	})
}
