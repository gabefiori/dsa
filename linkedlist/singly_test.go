package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingly(t *testing.T) {
	t.Run("Prepend", func(t *testing.T) {
		ll := NewSingly[int]()

		ll.Prepend(30)
		ll.Prepend(20)
		ll.Prepend(10)

		assert.Equal(t, 10, ll.Head.Value)
		assert.Equal(t, 20, ll.Head.Next.Value)
		assert.Equal(t, 30, ll.Head.Next.Next.Value)
	})

	t.Run("Append", func(t *testing.T) {
		ll := NewSingly[int]()

		ll.Append(10)
		ll.Append(20)
		ll.Append(30)

		assert.Equal(t, 10, ll.Head.Value)
		assert.Equal(t, 20, ll.Head.Next.Value)
		assert.Equal(t, 30, ll.Head.Next.Next.Value)
		assert.Equal(t, ll.Length(), 3)
	})

	t.Run("InsertAt", func(t *testing.T) {
		ll := NewSingly[int]()

		ll.Append(10)
		ll.Append(20)
		ll.Append(30)

		assert.NoError(t, ll.InsertAt(25, 2))
		assert.Equal(t, 25, ll.Head.Next.Next.Value)
		assert.Equal(t, ll.Length(), 4)

		assert.NoError(t, ll.InsertAt(5, 0))
		assert.Equal(t, 5, ll.Head.Value)
		assert.Equal(t, 5, ll.Length())

		assert.NoError(t, ll.InsertAt(40, 5))
		assert.Equal(t, 6, ll.Length())

		curr := ll.Head

		for curr.Next != nil {
			curr = curr.Next
		}

		assert.Equal(t, 40, curr.Value)

		assert.Error(t, ll.InsertAt(100, 100))
	})

	t.Run("DeleteFirst", func(t *testing.T) {
		ll := NewSingly[int]()

		assert.Error(t, ll.DeleteFirst())

		ll.Append(10)

		assert.NoError(t, ll.DeleteFirst())
		assert.Nil(t, ll.Head)
		assert.Equal(t, 0, ll.Length())

		ll.Append(10)
		ll.Append(20)

		assert.NoError(t, ll.DeleteFirst())
		assert.Equal(t, 20, ll.Head.Value)
		assert.Equal(t, 1, ll.Length())
	})

	t.Run("DeleteLast", func(t *testing.T) {
		ll := NewSingly[int]()

		assert.Error(t, ll.DeleteLast())

		ll.Append(10)

		assert.NoError(t, ll.DeleteLast())
		assert.Nil(t, ll.Head)
		assert.Equal(t, ll.Length(), 0)

		ll.Append(10)
		ll.Append(20)

		assert.NoError(t, ll.DeleteLast())
		assert.Equal(t, 10, ll.Head.Value)
		assert.Equal(t, 1, ll.Length())
	})

	t.Run("DeleteAt", func(t *testing.T) {
		ll := NewSingly[int]()

		assert.Error(t, ll.DeleteAt(0))

		ll.Append(10)

		assert.NoError(t, ll.DeleteAt(0))
		assert.Equal(t, ll.Length(), 0)

		ll.Append(10)
		ll.Append(20)

		assert.NoError(t, ll.DeleteAt(0))
		assert.Equal(t, 20, ll.Head.Value)
		assert.Equal(t, 1, ll.Length())

		ll.Append(30)
		ll.Append(40)

		assert.NoError(t, ll.DeleteAt(2))
		assert.Equal(t, 20, ll.Head.Value)
		assert.Equal(t, 30, ll.Head.Next.Value)
		assert.Equal(t, 2, ll.Length())
	})

	t.Run("Search", func(t *testing.T) {
		ll := NewSingly[int]()

		ll.Append(10)
		ll.Append(20)
		ll.Append(30)

		found, err := ll.Search(40)

		assert.Error(t, err)
		assert.Nil(t, found)

		found, err = ll.Search(30)

		assert.NoError(t, err)
		assert.Equal(t, found.Value, ll.Head.Next.Next.Value)
	})
}
