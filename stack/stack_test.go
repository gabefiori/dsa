package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	t.Run("Push", func(t *testing.T) {
		s := New[int]()

		s.Push(10)
		s.Push(20)
		s.Push(30)

		assert.Equal(t, 10, s.elems[0])
		assert.Equal(t, 20, s.elems[1])
		assert.Equal(t, 30, s.elems[2])
		assert.Equal(t, 3, s.Len())
	})

	t.Run("IsEmpty", func(t *testing.T) {
		s := New[int]()

		assert.True(t, s.IsEmpty())

		s.Push(10)
		assert.False(t, s.IsEmpty())
	})

	t.Run("Peek", func(t *testing.T) {
		s := New[int]()

		s.Push(10)
		s.Push(20)
		s.Push(30)

		elem, err := s.Peek()

		assert.NoError(t, err)
		assert.Equal(t, 30, elem)
	})

	t.Run("Pop", func(t *testing.T) {
		s := New[int]()

		s.Push(10)
		s.Push(20)
		s.Push(30)

		assert.Equal(t, 3, s.Len())

		elem, err := s.Pop()

		assert.NoError(t, err)
		assert.Equal(t, 30, elem)
		assert.Equal(t, 2, s.Len())

		elem, err = s.Pop()

		assert.NoError(t, err)
		assert.Equal(t, 20, elem)
		assert.Equal(t, 1, s.Len())

		elem, err = s.Pop()

		assert.NoError(t, err)
		assert.Equal(t, 10, elem)
		assert.Equal(t, 0, s.Len())

		_, err = s.Pop()
		assert.Error(t, err)
	})
}
